package passhash

// PasswordHasher defines the interface of a password hasher
type PasswordHasher interface {
	// Hash salts and hashes the given password returning the resulting hash
	Hash(password string) (string, error)

	// Compare returns true if the given password corresponds to
	// the given hash, otherwise returns false
	Compare(password, hash string) bool
}
