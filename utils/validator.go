package utils

import "strings"

func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}

	return strings.Contains(email, "@") && strings.Contains(email, ".")
}