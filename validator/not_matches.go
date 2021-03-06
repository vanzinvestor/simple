package validator

import "regexp"

// NotMatches If Matched, then return json error
func (v *validator) NotMatches(regex, str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	isMatched, _ := regexp.MatchString(regex, str)

	if isMatched {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
