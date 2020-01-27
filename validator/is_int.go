package validator

import (
	"reflect"
)

// IsInt If Not int, then show error message
func (v *validator) IsInt(value int, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	x := reflect.TypeOf(value).Kind()

	if x != reflect.Int {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
