package main

import (
	"core/pkg/config"
	"core/pkg/log"
	"core/pkg/server"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func main() {
	NewService()
}

func NewService() {
	var rootCmd = &cobra.Command{
		Version: "0.0.1",
		Use:     "myspider [command]",
		Short:   "MySpider CLI",
		Long:    `My spider web service`,
		Args:    cobra.MinimumNArgs(1),
	}

	var startServerCmd = &cobra.Command{
		Use:   "start ",
		Short: "start a web server to handle HTTP requests",
		Long:  `start a web server to handle HTTP requests and perform standard OIDC flow to interact with IAM`,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd)
		},
	}

	// 配置文件的绝对路径
	startServerCmd.Flags().StringP("configFile", "c", "", "the absolute path of config file that in yaml format")

	rootCmd.AddCommand(startServerCmd)

	if err := rootCmd.Execute(); err != nil {
		printCmdErr(err)
		os.Exit(1)
	}

}

func run(cmd *cobra.Command) {
	// 获取配置文件路径，加载并解析
	cfgPath, err := cmd.Flags().GetString("configFile")
	if err != nil {
		printCmdErr(err)
		return
	}
	if len(cfgPath) != 0 && !fileutil.IsExist(cfgPath) {
		printCmdErr(errors.New("The config file doesn't exist :" + cfgPath))
		return
	}

	PrintBanner()

	// 加载配置文件
	err = config.LoadConfig(cfgPath)
	if err == nil {
		cfg := config.GetConfig()

		if err = cfg.Validate(); err != nil {
			printCmdErr(err)
			return
		}

		if err = cfg.Complete(); err != nil {
			printCmdErr(err)
			return
		}

		// 初始化Log
		log.SetupLog(cfg)

		// 初始化web服务
		engine := server.Start()

		// 绑定地址，启动
		bindAddr := fmt.Sprintf("%v", cfg.BindAddress)
		if err := engine.Run(bindAddr); err != nil {
			log.Log().Error("Server fails to start", zap.Error(err))
		}
	}
}

// 在控制台输出出错信息
func printCmdErr(err error) {
	_, err = fmt.Fprintf(os.Stderr, "Error: '%s' \n", err)
	if err != nil {
		panic(err)
	}
}

func PrintBanner() {
	fmt.Println("My Spider v0.0.1")
}
