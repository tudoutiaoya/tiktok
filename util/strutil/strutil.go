package strutil

import "regexp"

// StringStrip 去除字符串空格
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	reg := regexp.MustCompile(`[\s\p{Zs}]{1,}`)
	return reg.ReplaceAllString(input, "")
}
