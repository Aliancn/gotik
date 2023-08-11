package common

// 100xx

const ErrCodeOK = 0
const ErrCodeInvalidArgs = 10000
const ErrcodePermissionDenied = 10001
const ErrcodeInternalError = 10003

func GetStatusMessage(code int) string {
	switch code {
	case ErrCodeOK:
		return "ok"
	case ErrcodePermissionDenied:
		return "permission denied"
	case ErrCodeInvalidArgs:
		return "invalid arguments"
	case ErrcodeInternalError:
		return "internal error, try again"
	}
	panic("unreachable")
}
