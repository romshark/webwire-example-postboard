package passhash

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// BcryptPasswordHasher implements the PasswordHasher interface using bcrypt
type BcryptPasswordHasher struct{}

// Hash salts and hashes the given password returning the resulting hash
func (h BcryptPasswordHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", errors.Wrap(err, "couldn't generate password hash")
	}
	return string(hash), nil
}

// Compare returns true if the given password corresponds to the given hash,
// otherwise returns false
func (h BcryptPasswordHasher) Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
