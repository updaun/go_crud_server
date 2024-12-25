package errors

import "fmt"

const (
	NotFoundUser = iota
)

var errMessages = map[int64]string{
	NotFoundUser: "not Found User",
}

func Errorf(code int64, args ...interface{}) error {
	if message, ok := errMessages[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
