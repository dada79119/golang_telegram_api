package util

import (
	"time"
	"strconv"
	"fmt"
)

const (
	IntervalSecond = "second"
	IntervalMinute = "minute"
	IntervalHour = "hour"
)

/*
函式功能：給予一個時間字串，計算與當下時間的間隔
    輸入：時間字串
    輸出：時間間隔分鐘
	其他：
*/
func TimeInterval(actionTime string) (interval float64) {
	local, _ := time.LoadLocation("Asia/Taipei") //修改成台北時間
	dateActTime, _ := time.ParseInLocation("20060102150405", actionTime, local)
	now := TimeNow()
	return now.Sub(dateActTime).Minutes()
}

/*
函式功能：取得當下的台北時間
    輸入：
    輸出：當下的台北時間
	其他：
*/
func TimeNow() time.Time {
	now := time.Now()
	local, _ := time.LoadLocation("Asia/Taipei") //修改成台北時間
	timeNow := now.In(local)
	return timeNow
}

// unix time string convert to time
func UnixTimeToTaipeiTime(unixTime string) (time.Time) {
	i, _ := strconv.ParseInt(unixTime, 10, 64)
	tempTime := time.Unix(i, 0)
	local, _ := time.LoadLocation("Asia/Taipei") //修改成台北時間
	taipeiTime := tempTime.In(local)
	return taipeiTime
}

const (
	FewSecondAgo = "幾秒前"
	FewMinuteAgo = "分鐘前"
	FewHourAge = "小時前"
	FewDayAge = "天前"
)

const (
	OneSecond = 1
	OneMinute = OneSecond*60
	OneHour = OneMinute*60
	OneDay = OneHour*24
	OneWeek = OneDay*7
)

/*
函式功能：取得 輸入時間 與 當下時間(台北時間) 的時間間隔 字眼
    輸入：
    輸出：當下的台北時間
	其他：
*/
func GetDateTimeDifferenceString(inputTime string) string {
	local, _ := time.LoadLocation("Asia/Taipei") //修改成台北時間
	inputTimeTaipei, _ := time.ParseInLocation("2006-01-02 15:04:05", inputTime, local)
	now := TimeNow()

	inputTimeDay := inputTimeTaipei.Year()
	currentTimeDay := now.Year()

	if inputTimeDay < currentTimeDay{
		// TODO：去年的時間，回傳： 2006年1月2日 的格式
		return  fmt.Sprintf("%d年%d月%d日", inputTimeTaipei.Year(),inputTimeTaipei.Month(),inputTimeTaipei.Day())
	}else if inputTimeDay == currentTimeDay{
		// TODO：同一年的時間，進行時間間隔判斷
		secondDiff := now.Sub(inputTimeTaipei).Seconds()	// 取得秒數差

		if secondDiff < OneMinute{
			// 小於 1分鐘，回傳幾秒前
			return FewSecondAgo
		}else if secondDiff < OneHour{
			// 大於1分鐘 且 小於 1小時，回傳 N分鐘前
			return fmt.Sprintf("%d%s",int(secondDiff/OneMinute), FewMinuteAgo)
		}else if secondDiff < OneDay{
			// 大於1小時 且 小於 24小時，回傳 N小時 前
			return fmt.Sprintf("%d%s",int(secondDiff/OneHour), FewHourAge)
		}else if secondDiff < OneWeek{
			// 大於1天 且 小於 7天，回傳 N天前，最多到7天前
			return fmt.Sprintf("%d%s",int(secondDiff/OneDay), FewDayAge)
		}else{
			// 大於等於7天，回傳 M月N日 的格式就好
			return  fmt.Sprintf("%d月%d日", inputTimeTaipei.Month(),inputTimeTaipei.Day())
		}
	}else{
		// TODO：未來的時間，回傳： 2006年1月2日  的格式
		return  fmt.Sprintf("%d年%d月%d日", inputTimeTaipei.Year(),inputTimeTaipei.Month(),inputTimeTaipei.Day())
	}
}
