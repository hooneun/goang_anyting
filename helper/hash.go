package helper

import "golang.org/x/crypto/bcrypt"

// MakeHash make a Hash!
func MakeHash(str string) (s string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHash !
func CompareHash(hash, s string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))

	if err != nil {
		return false, err
	}

	return true, nil
}
