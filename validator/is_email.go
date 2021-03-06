package validator

import (
	"github.com/asaskevich/govalidator"
)

// IsEmail If Not valid email, then return json error
func (v *validator) IsEmail(email, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if !govalidator.IsEmail(email) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
