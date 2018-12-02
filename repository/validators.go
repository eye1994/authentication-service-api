package repository

import (
	"regexp"
)

// ValidationError map where the key is the field and value is a slice of error messages
type ValidationError map[string][]string

// IsEmpty with(field string) return true if the field has 0 characters
func IsEmpty(field string) bool {
	return len(field) == 0
}

// IsEmail with(field string) return true if the field is not a valid email format
func IsEmail(field string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(field)
}
