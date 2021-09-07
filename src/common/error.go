package common

var (
	Success         = 0
	ErrInnerFault   = -76001
	ErrData         = -76002
	ErrHidataUpload = -76003
)

var ErrMap = map[int]string{
	Success:         "ok",
	ErrInnerFault:   "server busy",
	ErrData:         "err data, please check input parameters",
	ErrHidataUpload: "ErrHidataUpload",
}

var ErrKey = map[int]string{
	Success:         "Success",
	ErrInnerFault:   "ErrInnerFault",
	ErrData:         "ErrData",
	ErrHidataUpload: "ErrHidataUpload",
}
