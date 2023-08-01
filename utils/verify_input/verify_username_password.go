package verifyinput

import (
	"unicode"
)

// username的格式, 4~32个字符, 只能包含数字和字母
func IsUsernameValid(uname string) bool {
	if len(uname) > 32 || len(uname) < 3 {
		return false
	}

	for _, v := range uname {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

// password格式, 6~32个字符, 只能包含数字和字母
func IsPasswordValid(pword string) bool {
	if len(pword) > 32 || len(pword) < 6 {
		return false
	}

	for _, v := range pword {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}
