package u

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will generate a hashed password for us based on the user's input.
func HashPassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// VerifyPassword will check if passwords are matched.
func VerifyPassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}
