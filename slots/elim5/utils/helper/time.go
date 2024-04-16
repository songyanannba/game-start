package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// BetweenHourStr 两组 开始 结束 时间字符串判断是否交叉
func BetweenHourStr(s, e, ss1, ee1 string) bool {
	var (
		a1 = TimeSharingToInt(s) // 5
		a2 = TimeSharingToInt(e) // 1

		b1 = TimeSharingToInt(ss1) // 6
		b2 = TimeSharingToInt(ee1) // 2
	)
	arr1 := Hour24Parse(a1, a2)
	arr2 := Hour24Parse(b1, b2)
	return RangeOverlap(arr1, arr2)
}

func Hour24Parse(s1, s2 int) (arr []int) {
	hour24 := 24 * 60
	if s1 > s2 {
		arr = []int{s1, hour24 + s2}
	} else {
		arr = []int{s1, s2}
	}
	return arr
}

func TimeSharingToInt(s string) (i int) {
	arr := strings.Split(s, ":")
	h, _ := strconv.Atoi(SliceVal(arr, 0))
	m, _ := strconv.Atoi(SliceVal(arr, 1))
	i = h*60 + m
	return
}

func BetweenHourOrMinute(t time.Time, sTime, eTime string) bool {
	var (
		hour   = t.Hour()
		minute = t.Minute()
		stime  = strings.Split(sTime, ":")
		etime  = strings.Split(eTime, ":")
	)
	sHour, _ := strconv.Atoi(SliceVal(stime, 0))
	sMinute, _ := strconv.Atoi(SliceVal(stime, 1))
	eHour, _ := strconv.Atoi(SliceVal(etime, 0))
	eMinute, _ := strconv.Atoi(SliceVal(etime, 1))
	ok := false
	if GtTime(sHour, sMinute, eHour, eMinute) {
		if GtTime(hour, minute, sHour, sHour) || LtTime(hour, minute, eHour, eHour) {
			ok = true
		}
	} else {
		if GtTime(hour, minute, sHour, sHour) && LtTime(hour, minute, eHour, eHour) {
			ok = true
		}
	}
	return ok
}

func GtTime(h, m, sh, sm int) bool {
	ok := false
	if h == sh {
		if m > sm {
			ok = true
		}
		return ok
	}
	if h > sh {
		ok = true
	}
	return ok
}

func LtTime(h, m, sh, sm int) bool {
	ok := false
	if h == sh {
		if m < sm {
			ok = true
		}
		return ok
	}
	if h < sh {
		ok = true
	}
	return ok
}

func IndiaTimeZone() *time.Location {
	return time.FixedZone("CST", 5.5*3600)
}

func PRCTimeZone() *time.Location {
	return time.FixedZone("CST", 8*3600)
}

// GetDateByTime 获取时间的日期
func GetDateByTime(date time.Time, loc ...*time.Location) time.Time {
	// 获取时间戳
	if date.IsZero() {
		date = time.Now()
	}
	zone := time.Local
	if len(loc) > 0 {
		zone = loc[0]
	}
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, zone)
	return date
}

// GetDateByMonth 获取时间的月份
func GetDateByMonth(date time.Time, loc ...*time.Location) time.Time {
	if date.IsZero() {
		date = time.Now()
	}
	zone := time.Local
	if len(loc) > 0 {
		zone = loc[0]
	}
	date = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, zone)
	return date
}

// GetDateDiff 计算两个日期相隔的天数
func GetDateDiff(start, end time.Time) int {
	return int((end.Unix() - start.Unix()) / (24 * 60 * 60))
}

// IntDate 获取日期的整数格式 20201201
func IntDate(t time.Time) uint {
	i, _ := strconv.Atoi(t.Format("20060102"))
	return uint(i)
}

// DateInt converts an integer in the format (YYYYMMDD) to a time.Time.
func DateInt(i uint) time.Time {
	dateStr := fmt.Sprintf("%08d", i)
	t, _ := time.Parse("20060102", dateStr)
	return t
}

func BeforeDate(i uint, typ uint8) uint {
	t := DateInt(i)
	switch typ {
	case 3:
		return IntDate(t.AddDate(0, 0, -1))
	case 2:
		return IntDate(t.AddDate(0, -1, 0))
	default:
		return IntDate(t.AddDate(-1, 0, 0))
	}
}

func DateAdd(date uint, years, months, days int) uint {
	t := DateInt(date)
	return IntDate(t.AddDate(years, months, days))
}

// GetMonthDateInterval 获取月份的日期区间 比如当前为2020-01-10 则返回20201201,20210101
func GetMonthDateInterval(t time.Time) (uint, uint) {
	s, _ := strconv.Atoi(t.Format("200601") + "01")
	e, _ := strconv.Atoi(t.AddDate(0, 1, 0).Format("200601") + "01")
	return uint(s), uint(e)
}

func GetDayRange(t time.Time) (time.Time, time.Time) {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	s := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	e := time.Date(year, month, day, 23, 59, 59, 0, time.Local)
	return s, e
}

func CheckBetweenTime(betweenTime []time.Time) (err error) {
	if len(betweenTime) != 2 {
		return errors.New("the time range is incorrect")
	}
	startTime := betweenTime[0]
	endTime := betweenTime[1]
	sub := endTime.Sub(startTime)
	if sub > 31*24*time.Hour || sub < 0 {
		return errors.New("the time span cannot exceed 31 days")
	}
	return
}
