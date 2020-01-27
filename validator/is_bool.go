package validator

import (
	"reflect"
)

// IsBool If Not boolean, then show error message
func (v *validator) IsBool(value bool, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	x := reflect.TypeOf(value).Kind()

	if x != reflect.Bool {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
