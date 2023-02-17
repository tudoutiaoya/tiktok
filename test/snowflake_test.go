package test

import (
	"fmt"
	"testing"
	"tiktok/middleware/msnowflake"
)

func TestSnow(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(msnowflake.GenerateID.NextVal())
	}
}
