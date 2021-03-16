package services

import "golang.org/x/crypto/bcrypt"

// AuthService models the struct containing all functions relating to Authorization and Authentication
// in the application.
type AuthService struct{}

// NewAuthService creates and returns a pointer to a new instance of AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// HashPassword hashes and salts a supplied password and returns the hash.
func (a *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(hash), err
	}

	return string(hash), nil
}
