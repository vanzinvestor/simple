package validator

func (v *validator) MustBeLongerThan(value string, min int, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if value == "" {
		return true
	}

	if len(value) < min {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
