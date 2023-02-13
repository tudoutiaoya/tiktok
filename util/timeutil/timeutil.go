package timeutil

import "time"

func Time2mm_dd(time time.Time) string {
	return time.Format("01-02")
}
