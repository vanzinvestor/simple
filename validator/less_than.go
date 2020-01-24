package validator

func (v *validator) MustBeLessThan(value string, max int, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if value == "" {
		return true
	}

	if len(value) > max {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
