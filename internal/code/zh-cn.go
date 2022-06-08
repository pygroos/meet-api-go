package code

var MsgFlags = map[int]string {
	SUCCESS:       "成功",
	ERROR_CLOSE:   "关闭评分失败",
	ERROR_DELETE:   "删除评分失败",
	ERROR_SUBMIT: 	"提交评分失败",
	ERROR_UNkNOWN: "未知错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR_UNkNOWN]
}