package strutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// StringStrip 去除字符串空格
func TestStringStrip(t *testing.T) {
	a := assert.New(t)
	str := "  1  a  1 "
	strip := StringStrip(str)
	a.Equal(strip, "1a1")
}
