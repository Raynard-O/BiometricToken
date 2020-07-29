package lib

import "golang.org/x/crypto/bcrypt"

var PasswordIncoreect = "Password Entered Incorrect"



func GenerateHashFromPassword(password string) string {
	hash , err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		panic("error generating hash from password")
	}
	return string(hash)
}

func CompareHashWithPassword(hash , password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash) ,[]byte(password))
	return err == nil
}
