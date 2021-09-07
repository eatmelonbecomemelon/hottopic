package weibo

import (
	cm "common"
	"gojieba"
	"mgo/bson"
	"mongoplus"
)

type WordInfo struct {
	Word  string
	Month string
	Count int
}

func sentenceParse(sentence string) (words []string) {
	x := gojieba.NewJieba()
	return x.CutAll(sentence)
}

func (w *WordInfo) UpdateCurrentCount() {
	var currentCnt int
	filter := bson.M{"word": w.Word, "month": w.Month}
	data, err := mongoplus.MgoEntry.QueryOne(hotwordCol, filter)
	if err != nil {
		cm.Error(err.Error())
		return
	}
	if len(data) == 1 {
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
