package time

import (
	"time"
)

func Int64BeginTime(date int64) time.Time {
	beginTime := time.Unix(date, 0)
	return BeginTime(beginTime)
}

func Int64EndTime(date int64) time.Time {
	endTime := time.Unix(date, 0)
	return EndTime(endTime)
}

func BeginTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
}

func DatetimeStrToTimeBegin(t string) time.Time {
	var (
		layout = "2006-01-02 15:04:05"
	)

	dateTime, _ := time.Parse(layout, t)

	return BeginTime(dateTime)
}

func DatetimeStrToTimeEnd(t string) time.Time {
	var (
		layout = "2006-01-02 15:04:05"
	)

	dateTime, _ := time.Parse(layout, t)
	return EndTime(dateTime)
}

func DateStrToTimeBegin(t string) time.Time {
	dateTime, _ := time.Parse(time.DateOnly, t)

	return BeginTime(dateTime)
}

func DateStrToTimeEnd(t string) time.Time {
	dateTime, _ := time.Parse(time.DateOnly, t)

	return EndTime(dateTime)
}

func ToDatetimeStr(date time.Time) string {
	return date.Format(time.DateTime)
}

func StrToDatetime(date string) time.Time {
	var (
		layout = "2006-01-02 15:04:05"
	)

	dateTime, _ := time.Parse(layout, date)
	return dateTime
}

func UnixToDatetimePtr(date int64) *time.Time {
	dateTime := time.Unix(date, 0)
	return &dateTime
}

func DatetimeStrToDatetimePtr(date string) *time.Time {
	var (
		layout = "2006-01-02 15:04:05"
	)

	dateTime, _ := time.Parse(layout, date)
	return &dateTime
}

func DateToDateStr(date time.Time) string {
	return date.Format(time.DateOnly)
}

func DateStrToDate(date string) time.Time {
	var (
		layout = "2006-01-02"
	)

	dateTime, _ := time.Parse(layout, date)
	return dateTime
}
