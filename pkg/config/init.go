package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/pflag"
)

const (
	DefaultPort = 0
	DefaultMode = ""
)

// 启动参数
var (
	configFile  = pflag.StringP("config", "f", "config/setting.yml", "config file")
	showVersion = pflag.BoolP("version", "v", false, "Print version information.")
)

// 环境配置参数

func ParseFlag() {
	pflag.Parse()

	if *showVersion {
		fmt.Printf("Version %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
		os.Exit(0)
	}

	fmt.Println("starting init pflag info")
}

func initConfig() {

}
