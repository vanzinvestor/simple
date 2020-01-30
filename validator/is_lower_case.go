package validator

import "github.com/asaskevich/govalidator"

// IsLower If character Not lowercase, then return json error
func (v *validator) IsLowerCase(str string, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if !govalidator.IsLowerCase(str) {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
