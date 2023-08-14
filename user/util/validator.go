package util

func ValidateParam(s string) bool {
	if len(s) > 32 || len(s) < 0 {
		return false
	}
	return true
}

func ValidateParams(ss []string) bool {
	for i := 0; i < len(ss); i++ {
		if !ValidateParam(ss[i]) {
			return false
		}
	}
	return true
}
