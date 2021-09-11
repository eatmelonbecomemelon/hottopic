package router

import (
	cm "common"
	"github.com/gin-gonic/gin"
	"net/http"
	"proxyfront"
	"weibo"
)

//Init routers

func InitRouter() *gin.Engine {
	gin.SetMode("debug")
	r := gin.Default()
	r.POST("/hottopic/common", handleCommonReq)
	r.StaticFS("/hottopic/file", http.Dir("."))

	return r
}

func handleCommonReq(c *gin.Context) {
	var (
		ret    int
		detail string
		result interface{}
	)
	defer func() {
		c.String(http.StatusOK, proxyfront.FormatResp(ret, result, cm.ErrMap, detail))
		return
	}()

	var reqData struct {
		Data map[string]interface{} `json:"data"`
		Ope  string                 `json:"ope"`
	}
	err := c.ShouldBindJSON(&reqData)

	if err != nil {
		ret = cm.ErrData
		detail = err.Error()
		return
	}
	switch reqData.Ope {
	case weibo.OpeGetHotWord, weibo.OpeGetHottopic:
		result, ret, detail = weibo.Handler(reqData.Data, reqData.Ope)
	default:
		cm.Error("not supported ope:", reqData.Ope)
		return
	}
}
