package log

import (
	"core/pkg/config"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevelMap = map[string]zapcore.Level{
	"DEBUG": zapcore.DebugLevel,
	"INFO":  zapcore.InfoLevel,
	"WARN":  zapcore.WarnLevel,
	"ERROR": zapcore.ErrorLevel,
}

// zap.SugaredLogger 就是对 zap.Logger 进行了封装，提供了一些高级的语法糖特性，
// 如支持使用类似 fmt.Printf 的形式进行格式化输出
var sugLogger *zap.SugaredLogger
var logger *zap.Logger

func SetupLog(cfg *config.SysConfig) {
	var logSetting = cfg.LogSetting
	level, ok := logLevelMap[logSetting.LogLevel]
	if !ok {
		panic("the log_level is invalid, it only supports: DEBUG,INFO,WARN,ERROR.")
	}

	atomicLevel := zap.NewAtomicLevelAt(level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		// EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 日志轮转
	writer := &lumberjack.Logger{
		// 日志名称
		Filename: logSetting.FileName,
		// 日志大小限制，单位MB
		MaxSize: logSetting.MaxSizeInMB,
		// 历史日志文件保留天数
		MaxAge: logSetting.MaxAgeInDay,
		// 最大保留历史日志数量
		MaxBackups: logSetting.MaxBackups,
		// 本地时区
		LocalTime: true,
		// 历史日志文件压缩
		Compress: logSetting.Compress,
	}

	syncers := []zapcore.WriteSyncer{zapcore.AddSync(writer)}
	if logSetting.OutputConsole {
		// 同时输出到控制台
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(syncers...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	options := []zap.Option{zap.AddCaller()}
	options = append(options, zap.Development())

	logger = zap.New(zapCore, options...)
	sugLogger = logger.Sugar()

	// 替换全局logger
	zap.ReplaceGlobals(logger)
}

func Logger() *zap.Logger {
	return logger
}

func SugaredLogger() *zap.SugaredLogger {
	return sugLogger
}
