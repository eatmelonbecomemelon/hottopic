package httpplus

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
	cm "common"
)

func getHttpClient() (client *http.Client) {
	transport := http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{
		Timeout: 15 * time.Second,
		Transport: &transport,
	}
	return
}

func HttpRequest(method string, backUrl string, header map[string]string, body []byte) (respbyte []byte, err error, costMs int64) {
	var timeStart int64
	req, _ := http.NewRequest(method, backUrl, bytes.NewReader(body))
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := getHttpClient()
	timeStart = time.Now().UnixNano()
	defer func() {
		costMs = (time.Now().UnixNano() - timeStart) / 1e6
	}()
	resp, err := client.Do(req)
	if err != nil {
		costMs = (time.Now().UnixNano() - timeStart) / 1e6
		cm.Error(err.Error(), "costTime:", costMs)
		return
	}
	respbyte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		cm.Error(err.Error())
		return
	}
	return
}