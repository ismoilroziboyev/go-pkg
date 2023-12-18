package validation

import "net/mail"

func IsValidEmail(v string) bool {
	_, err := mail.ParseAddress(v)
	return err == nil
}
