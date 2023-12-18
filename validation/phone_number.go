package validation

import "strings"

func IsValidPhone(phone string) bool {
	if len(phone) != 12 {
		return false
	}

	if !strings.HasPrefix(phone, "998") {
		return false
	}

	return true
}
