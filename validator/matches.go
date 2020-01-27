package validator

import "regexp"

// Matches If Not Matched, then show error message
func (v *validator) Matches(regex, value, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	isMatched, _ := regexp.MatchString(regex, value)

	if !isMatched {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
