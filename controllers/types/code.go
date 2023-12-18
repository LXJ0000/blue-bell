package types

type Code int64

const (
	CodeSUCCESS Code = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[Code]string{
	CodeSUCCESS:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
}

func getMsg(code Code) string {
	if msg, ok := codeMsgMap[code]; ok {
		return msg
	}
	return codeMsgMap[CodeServerBusy]
}
