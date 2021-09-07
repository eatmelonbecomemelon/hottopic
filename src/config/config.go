package cf

import (
	cm "common"
	"flag"
	"fmt"
	"os"
	"toml"
)

var SysConfig struct {
	MongoDB map[string]struct {
		IP             string
		Port           string
		User           string
		PassWord       string
		DBName         string
		MongoPoolLimit int
	}
	WeiboHotRunHour []string
	LogLevel        string
}

func LoadConfig() {
	showVersion := flag.Bool("v", false, "show version")
	flag.PrintDefaults()
	flag.Parse()
	if *showVersion {
		os.Exit(0)
	}
	configPath := flag.String("f", "./output/hottopic.toml", "configfile")
	cm.Debug("cfg path :", *configPath)
	if *configPath == "" {
		*configPath = "./output/vocvtdock.toml"
	}

	if _, err := toml.DecodeFile(*configPath, &SysConfig); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("%+v\n", SysConfig)
	cm.SetLogLevel(SysConfig.LogLevel)
}
