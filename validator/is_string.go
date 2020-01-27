package validator

import (
	"reflect"
)

// IsString If Not string, then show error message
func (v *validator) IsString(value string, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	x := reflect.TypeOf(value).Kind()

	if x != reflect.String {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
