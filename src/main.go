package main

import (
	cm "common"
	cf "config"
	"mongoplus"
	"robfig/cron"
	"weibo"
)

var mgoEntity mongoplus.MgoDB

func main() {
	cf.LoadConfig()
	mongoplus.MongoInit()
	runCronTabJob()
	select {}
}

func runCronTabJob() {
	cm.Info("START")
	crontab := cron.New()
	for _, one := range cf.SysConfig.WeiboHotRunHour {
		crontab.AddFunc("00 "+one+" * * *", weibo.GetWeiboHotTopic)
	}
	crontab.Start()
}
