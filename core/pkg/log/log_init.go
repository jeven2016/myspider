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
}

func Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	if ce := logger.Check(lvl, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if ce := logger.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Info(msg string, fields ...zap.Field) {
	if ce := logger.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if ce := logger.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if ce := logger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Logger() *zap.Logger {
	return logger
}
