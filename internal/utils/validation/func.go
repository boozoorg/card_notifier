package validation

import "strings"

// IsOrderTypeCorrect is a simple validation check, can be added known orders types to send correct message to client
func IsOrderTypeCorrect(ot string) bool {
	if ot == "" {
		return false
	}
	return true
}

// IsCardCorrect is a simple validation check, can be added correct card len checker
func IsCardCorrect(c string) bool {
	if c == "" || len(c) != 10 {
		return false
	}
	return true
}

// IsWebUrlCorrect is a simple validation check, can be added url contains of domain (like '.com', '.tj' and so on)
func IsWebUrlCorrect(wu string) bool {
	if wu == "" || !strings.HasPrefix(wu, "https://") {
		return false
	}
	return true
}
