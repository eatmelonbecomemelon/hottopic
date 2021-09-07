package weibo

import (
	cm "common"
	"encoding/json"
	"httpplus"
	"mongoplus"
	"time"
)

type weiboResp struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	NewsList []HotInfo `json:"newslist"`
}

type HotInfo struct {
	HotWord    string `json:"hotword"`
	HotWordNum string `json:"hotwordnum"`
	HotTag     string `json:"hottag"`
}
type HotInfoRecord struct {
	HotInfo
	Date string `json:"date"`
}

const dayLayout = "2006-01-02"
const weibohostspotCol = "weibohostspot"
const hotwordCol = "hotword"

func GetWeiboHotTopic() {
	getUrl := "http://api.tianapi.com/txapi/weibohot/index?key=26b1f3800f766b36270aaf7ec734cbac"
	resp, err, _ := httpplus.HttpRequest("GET", getUrl, nil, nil)
	if err != nil {
		cm.Error(err.Error())
		return
	}
	var respData weiboResp
	err = json.Unmarshal(resp, &respData)
	if err != nil {
		cm.Error(err.Error())
		return
	}
	date := time.Now().Format(dayLayout)
	var recodes []HotInfoRecord
	var record = HotInfoRecord{
		Date: date,
	}
	var allWords = make(map[string]WordInfo)
	var month = time.Now().Format("2006-01")
	for _, one := range respData.NewsList {
		record.HotInfo = one
		words := sentenceParse(record.HotWord)
		recodes = append(recodes, record)
		for _, one := range words {
			wordInfo, ok := allWords[one]
			if !ok {
				wordInfo.Word = one
				wordInfo.Month = month
			}
			wordInfo.Count++
		}
	}

	ret := mongoplus.InsertRecords(recodes, "weibohostspot")
	if ret != cm.Success {
		cm.Error(ret)
		return
	}
	for _, wordInfo := range allWords {
		wordInfo.UpdateCurrentCount()
	}
	cm.Info("Update records", len(recodes))
	cm.Info("Update word", len(recodes))

}
