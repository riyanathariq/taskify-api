package common

import "testing"

func TestHashPassword(t *testing.T) {
	password := "102938"
	t.Run("test", func(t *testing.T) {
		got, err := HashPassword(password)
		if err != nil {
			t.Errorf("HashPassword() error = %v", err)
			return
		} else {
			t.Logf("success, got: %s", got)
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	password := "102938"
	hash := "$2a$10$LdR4hZgxIWRi3/OM504BOOgVddYq0xPvq5rAJD.9zsXbX0dcycKOu"
	t.Run("test", func(t *testing.T) {
		t.Logf("success, got: %v", CheckPasswordHash(password, hash))

	})
}
