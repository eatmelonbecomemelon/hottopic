package main

import (
	cf "config"
	"mongoplus"
	"testing"
	"weibo"
)

func Test_GetWeiboHotTopic(t *testing.T) {
	cf.LoadConfig()
	mongoplus.MongoInit()
	weibo.GetWeiboHotTopic()
}
