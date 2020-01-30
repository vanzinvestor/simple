package validator

import (
	"github.com/asaskevich/govalidator"
)

// HasUpperCase If has Not uppercase, then return json error
func (v *validator) HasUpperCase(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if !govalidator.HasUpperCase(str) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
