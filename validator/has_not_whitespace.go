package validator

import (
	"github.com/asaskevich/govalidator"
)

// HasNotWhitespace If has whitespace, then return json error
func (v *validator) HasNotWhitespace(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if govalidator.HasWhitespace(str) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
