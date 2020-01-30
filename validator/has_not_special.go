package validator

import "regexp"

// HasNotSpecial If has special character, then return json error
func (v *validator) HasNotSpecial(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	isMatched, _ := regexp.MatchString(`[!~@#$%^&*(),+-/_.?":;{}|<>]`, str)

	if isMatched {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
