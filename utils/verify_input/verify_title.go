package verifyinput

import "unicode/utf8"

// 4 ~ 35个字符
func IsTitleValid(title string) bool {
	c := utf8.RuneCountInString(title)
	return c >= 4 && c <= 35
}
