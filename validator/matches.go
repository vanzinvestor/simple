package validator

import "regexp"

// Matches If Not Matched, then return json error
func (v *validator) Matches(regex, str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	isMatched, _ := regexp.MatchString(regex, str)

	if !isMatched {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
