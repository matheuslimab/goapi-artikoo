package helpers

import "regexp"

// Valida se um nome contem somente letras
func ValidateString(str string, reg string) bool {
	regex := regexp.MustCompile(reg)

	return regex.MatchString(str)
}
