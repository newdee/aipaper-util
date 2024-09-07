package auth

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Encrypt md5加密方式
func MD5Encrypt(password string) string {
	hash := md5.Sum([]byte(password))
	encryptedPassword := hex.EncodeToString(hash[:])
	return encryptedPassword
}
