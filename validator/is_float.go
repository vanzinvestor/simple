package validator

import (
	"reflect"
)

// IsFloat If Not float64, then show error message
func (v *validator) IsFloat(value float64, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	x := reflect.TypeOf(value).Kind()

	if x != reflect.Float64 {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
