package weibo

import "github.com/yanyiwu/gojieba"

func init() {
	gojiebaEntity = gojieba.NewJieba()
	gojiebaEntity.AddWord("张晚意")
	gojiebaEntity.AddWord("乔四美")
	gojiebaEntity.AddWord("乔家的儿女")
}
