package timeutil

import "time"

func Time2mm_dd(time time.Time) string {
	return time.Format("01-02")
}

func Time2standardTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
