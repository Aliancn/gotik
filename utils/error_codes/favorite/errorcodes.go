package favorite

// 104xx

const (
	ErrCodeOptionWrong     = 10400
	ErrCodeOptionTypeWrong = 10401
)

func GetStatusMessage(code int) string {
	switch code {
	case ErrCodeOptionWrong:
		return "something bad happened"
	case ErrCodeOptionTypeWrong:
		return "something bad happened"
	}
	panic("favorite option unsuccessful")
}
