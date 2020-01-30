package validator

import (
	"github.com/asaskevich/govalidator"
)

// IsLower If character Not uppercase, then return json error
func (v *validator) IsUpperCase(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if !govalidator.IsUpperCase(str) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
