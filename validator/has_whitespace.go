package validator

import (
	"github.com/asaskevich/govalidator"
)

// HasWhitespace If has Not whitespace, then return json error
func (v *validator) HasWhitespace(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if !govalidator.HasWhitespace(str) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
