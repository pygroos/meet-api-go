package support

import "time"

// GetDate 获取指定日期
// GetDate(-1) 昨天
// GetDate(1) 明天
func GetDate(flag int) string {
	current := time.Now()
	before := current.AddDate(0, 0, flag)
	return before.Format("2006-01-02")
}
