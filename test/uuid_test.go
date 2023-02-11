package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	str := uuid.New()

	fmt.Println(str)
}
