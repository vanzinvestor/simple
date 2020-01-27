package validator

import "unicode"

// IsLower If character Not lowercase, then show error message
func (v *validator) IsLower(value string, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	for _, r := range value {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			v.Errors["error"] = errMessage{message: errorMessage}.Error()
			return false
		}
	}

	return true
}
