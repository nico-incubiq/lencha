package utils

import "regexp"

func IsAlphaNumeric(str string) bool {
	return !regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString(str)
}

func IsEmail(str string) bool {
	return regexp.MustCompile(`(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`).MatchString(str)
}
