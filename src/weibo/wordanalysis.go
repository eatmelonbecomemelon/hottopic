package weibo

import (
	cm "common"
	"github.com/yanyiwu/gojieba"
	"mgo/bson"
	"mongoplus"
	"strings"
)

type WordInfo struct {
	Word  string
	Month string
	Count int
}

var gojiebaEntity *gojieba.Jieba

func sentenceParse(sentence string) (words []string) {
	use_hmm := true
	return gojiebaEntity.Cut(sentence, use_hmm)
}

func (w *WordInfo) UpdateCurrentCount() {
	var currentCnt int
	filter := bson.M{"word": w.Word, "month": w.Month}
	data, err := mongoplus.MgoEntry.QueryOne(hotwordCol, filter)
	if err != nil {
		cm.Error(err.Error())
		if !strings.Contains(err.Error(), "not found") {
			return
		}
	}
	if len(data) > 1 {
		currentCnt = cm.MustInt(data["count"])
	}
	w.Count = currentCnt + w.Count
	err = mongoplus.MgoEntry.Upsert(hotwordCol, filter, w)
	if err != nil {
		cm.Error(err.Error())
		return
	}
	return
}
