package validator

import "unicode"

// IsLower If character Not uppercase, then show error message
func (v *validator) IsUpper(value string, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	for _, r := range value {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			v.Errors["error"] = errMessage{message: errorMessage}.Error()
			return false
		}
	}

	return true
}
