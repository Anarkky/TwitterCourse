package db

import "golang.org/x/crypto/bcrypt"

/* */
func EncryptPassword(pass string) (string, error) {
	//cost is the number of encrypts to the password 2^cost ex: 2^8
	cost := 8
	//slice is a byte array without element size
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
