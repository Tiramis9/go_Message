package check

import (
	"time"
)

//如果操作时间为当天，则false,如果当前时间小于14点则false
func CheckTime(actTime int64) bool {
	integral := int64(28800) //	(8H)
	fourteen := int64(50400) //(14H)
	nowTime := time.Now().Unix()
	timeStr := time.Now().Format("2006-01-02")
	today, _ := time.Parse("2006-01-02", timeStr)
	zerotimes := today.Unix() - integral
	pmTow := zerotimes + fourteen // PM 2点(14H)
	if actTime >= zerotimes || nowTime <= pmTow {
		return false
	}
	return true
}
