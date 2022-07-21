package util

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

var options = &password.Options{
	SaltLen:      16,
	Iterations:   100,
	KeyLen:       32,
	HashFunction: sha512.New,
}

func EncryptPassword(originPassword string) string {
	salt, encodedPwd := password.Encode(originPassword, options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return newPassword
}

func VerifyPassword(encryptedPassword string, originPassword string) bool {
	passwordInfo := strings.Split(encryptedPassword, "$")
	check := password.Verify(originPassword, passwordInfo[2], passwordInfo[3], options)
	return check
}
