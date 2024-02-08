package validation

// IsOrderTypeCorrect is a simple validation check, can be added known orders types to send correct message to client
func IsOrderTypeCorrect(ot string) bool {
	if ot == "" {
		return false
	}
	return true
}

// IsCardCorrect is a simple validation check, can be added card len checker
func IsCardCorrect(c string) bool {
	if c == "" {
		return false
	}
	return true
}

// IsWebUrlCorrect is a simple validation check, can be added url scheme checker and contains of domain (like '.com', '.tj' and so on)
func IsWebUrlCorrect(wu string) bool {
	if wu == "" {
		return false
	}
	return true
}
