package validator

// IsLengthMax If value more than maximum, then return json error
func (v *validator) IsLengthMax(str string, max int, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if str == "" {
		return true
	}

	if len(str) > max {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
