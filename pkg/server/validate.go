package server

import (
	DB "github.com/onurcanari/kartaca_spa/pkg/db"
)

// ValidateUserPassword Validate the user.
func ValidateUserPassword(creds *Credentials) bool {
	pass, _ := DB.GetUserPassword(creds.Username)
	if pass == creds.Password {
		return true
	}
	return false
}
