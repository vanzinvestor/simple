package validator

// IsLengthMin If value less than minimum, then return json error
func (v *validator) IsLengthMin(str string, min int, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if str == "" {
		return true
	}

	if len(str) < min {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
