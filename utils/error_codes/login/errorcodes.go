package login

// 102xx

const ErrCodeUsernameOrPasswordWrong = 10000

func GetStatusMessage(code int) string {
	switch code {
	case ErrCodeUsernameOrPasswordWrong:
		return "username or password is wrong"
	}
	panic("unreachable")
}
