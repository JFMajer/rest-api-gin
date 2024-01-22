package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plaintext password and returns its bcrypt hash.
// The bcrypt algorithm is used for securely hashing passwords.
// The second argument to GenerateFromPassword is the cost of hashing,
// which determines how computationally expensive the hash calculation is.
func HashPassword(password string) (string, error) {
	// GenerateFromPassword hashes the password using bcrypt.
	// It returns the hashed password as a byte slice.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		// If an error occurred during hashing, return an empty string and the error.
		return "", err
	}
	// Convert the hashed password (bytes) to a string and return it.
	return string(bytes), nil
}

// VerifyPassword compares a hashed password with a plaintext password.
// It returns true if the passwords match, false otherwise.
func VerifyPassword(passFromDB string, passFromForm string) bool {
	// CompareHashAndPassword compares the bcrypt hashed password with its possible
	// plaintext equivalent. Returns nil on success, or an error on failure.
	err := bcrypt.CompareHashAndPassword([]byte(passFromDB), []byte(passFromForm))

	// If err is nil, the passwords match. Otherwise, they don't.
	return err == nil
}
