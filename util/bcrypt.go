package util

import (
	"golang.org/x/crypto/bcrypt"
)

func ENCRYPT_BCRYPT(password []byte, round int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 12)
}

func VERIFY_BCRYPT(compare, password []byte) error {
	return bcrypt.CompareHashAndPassword(compare, password)
}
