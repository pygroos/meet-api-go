package code

var MsgFlags = map[int]string{
	SUCCESS:          "成功",
	USER_LOGIN_ERROR: "登录错误",
	ERROR_UNkNOWN:    "未知错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR_UNkNOWN]
}
