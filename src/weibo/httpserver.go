package weibo

import (
	cm "common"
)

const (
	OpeGetHottopic = "gethottopic"
	OpeGetHotWord  = "gethotword"
)

func Handler(data map[string]interface{}, ope string) (result interface{}, ret int, detail string) {
	var (
		err error
	)
	defer func() {
		if err != nil {
			ret = cm.ErrInnerFault
			detail = err.Error()
		}
		return
	}()
	switch ope {
	case OpeGetHotWord:
		result, err = GetTopWord(cm.MustString(data["month"]))
	}
	return
}
