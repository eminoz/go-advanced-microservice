package encryption

import "golang.org/x/crypto/bcrypt"

type Encryption interface {
	GenerateHashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type UserEncryption struct {
}

func (e UserEncryption) GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (e UserEncryption) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
