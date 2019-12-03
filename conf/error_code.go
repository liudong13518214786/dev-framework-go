package conf

const (
	SUCCESS         = 100
	ERROR           = 500
	INVALID_PARAMS  = 400
	ValidationError = 401
)

var MsgContent = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "服务器异常",
	INVALID_PARAMS:  "参数错误",
	ValidationError: "认证失败",
}

func GetMessage(code int) string {
	value, ok := MsgContent[code]
	if ok {
		return value
	}
	return "未知错误"
}
