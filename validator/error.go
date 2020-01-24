package validator

import "fmt"

type errMessage struct {
	message string
}

func (e errMessage) Error() string {
	return fmt.Sprintf("%v", e.message)
}
