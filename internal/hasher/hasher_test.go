package hasher

import "testing"

func TestHashAndCheckHash(t *testing.T) {

	hasher := NewHasher()

	t.Run("Проверка валидности работы", func(t *testing.T) {
		password := "12345"
		res, err := hasher.Hash(password)
		if err != nil {
			t.Error(err)
		}

		if check := hasher.CheckHash(res, password); !check {
			t.Error("Полученный хэш невалиден")
		}
	})

}
