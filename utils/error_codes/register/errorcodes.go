package register

// 101xx

const ErrCodeUsernameOccupied = 10100

func GetStatusMessage(code int) string {
	switch code {
	case ErrCodeUsernameOccupied:
		return "username is occupied"
	}
	panic("unreachable")
}
