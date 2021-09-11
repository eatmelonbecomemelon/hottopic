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
	Index int    `json:"index"`
	Hour  string `json:"hour"`
}

const dayLayout = "2006-01-02"
const hourLayout = "2006-01-02_15"
const weibohostspotCol = "weibohostspot2"
const hotwordCol = "hotword2"

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
	hour := time.Now().Format(hourLayout)
	var recodes []HotInfoRecord
	var record = HotInfoRecord{
		Hour: hour,
	}
	var allWords = make(map[string]WordInfo)
	var month = time.Now().Format("2006-01")
	for i, one := range respData.NewsList {
		record.HotInfo = one
		record.Index = i
		words := sentenceParse(record.HotWord)
		recodes = append(recodes, record)
		for _, one := range words {
			wordInfo, ok := allWords[one]
			if !ok {
				wordInfo.Word = one
				wordInfo.Month = month
			}
			wordInfo.Count++
			allWords[one] = wordInfo
		}
	}

	ret := mongoplus.InsertRecords(recodes, weibohostspotCol)
	if ret != cm.Success {
		cm.Error(ret)
		return
	}
	for _, wordInfo := range allWords {
		wordInfo.UpdateCurrentCount()
	}
	cm.Info("Update records", len(recodes))
	cm.Info("Update word", len(allWords))

}
