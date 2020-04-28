package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateHash(str []byte)(string, error) {
	// Hash Password
	hash, err := bcrypt.GenerateFromPassword(str, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHash(str1 string, str2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(str1), []byte(str2))
	if err != nil {
		return err
	}
	return nil
}
