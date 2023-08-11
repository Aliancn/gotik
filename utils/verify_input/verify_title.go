package verifyinput

import "unicode/utf8"

// 1 ~ 35个字符
func IsTitleValid(title string) bool {
	c := utf8.RuneCountInString(title)
	return c >= 1 && c <= 35
}
