package user

import (
	"os"

	"golang.org/x/crypto/argon2"
)

var hashSalt = os.Getenv("PASSWORD_HASH_SALT")

// HashPassword hashes the given password using the Argon2 KDF.
func HashPassword(password string) []byte {
	return argon2.IDKey([]byte(password), []byte(hashSalt), 2, 19*1024, 1, 32)
}
