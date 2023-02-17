package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())

	for i := 0; i < 1000; i++ {
		n := rand.Intn(10) + 1
		fmt.Println(n)
	}
}
