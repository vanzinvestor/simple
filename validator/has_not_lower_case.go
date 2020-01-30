package validator

import (
	"github.com/asaskevich/govalidator"
)

// HasNotLowerCase If has lowercase, then return json error
func (v *validator) HasNotLowerCase(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if govalidator.HasLowerCase(str) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
