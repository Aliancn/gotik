package common

// 100xx

const ErrCodeOK = 0
const ErrCodeInvalidArgs = 10000
const ErrCodePermissionDenied = 10001
const ErrCodeInternalError = 10003
const ErrCodeNotLogin = 10004

func GetStatusMessage(code int) string {
	switch code {
	case ErrCodeOK:
		return "ok"
	case ErrCodePermissionDenied:
		return "permission denied"
	case ErrCodeInvalidArgs:
		return "invalid arguments"
	case ErrCodeInternalError:
		return "internal error, try again"
	case ErrCodeNotLogin:
		return "Not login, Please login and try again"
	}
	panic("unreachable")
}
