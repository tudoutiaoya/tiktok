package timeutil

import (
	"fmt"
	"testing"
	"time"
)

func TestTime2mm_dd(t *testing.T) {
	now := time.Now()
	fmt.Println(Time2mm_dd(now))
}
