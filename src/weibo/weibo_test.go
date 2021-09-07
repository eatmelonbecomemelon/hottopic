package weibo

import (
	cf "config"
	"mongoplus"
	"testing"
)

func Test_GetWeiboHotTopic(t *testing.T) {
	cf.LoadConfig()
	mongoplus.MongoInit()
}
