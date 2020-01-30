package validator

import "regexp"

// HasNotDigit If has digit, then return json error
func (v *validator) HasNotDigit(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	isMatched, _ := regexp.MatchString(`[\d]`, str)

	if isMatched {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
