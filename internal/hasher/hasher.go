package hasher

import "golang.org/x/crypto/bcrypt"

// BcryptHasher - структура для хэширования паролей имплементирует Hasher
type BcryptHasher struct {
}

func NewHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *BcryptHasher) CheckHash(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
