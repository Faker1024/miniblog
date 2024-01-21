package auth

import "golang.org/x/crypto/bcrypt"

// Encrypt 使用bcrypt加密纯文本
func Encrypt(source string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashBytes), err
}

// Compare 比较密文和明文是否相同
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}