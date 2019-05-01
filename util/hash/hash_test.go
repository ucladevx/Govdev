package hash

import (
	"testing"
)

func TestArgon2ID(t *testing.T) {

	argon := NewArgon2ID(
		64*1024,
		3,
		2,
		16,
		32,
	)
	password := "password123"
	fakePassword := "pa$$sword"

	encodedHash, err := argon.HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error hashing passwords: %s\n", err.Error())
	}

	match, err := argon.ComparePasswordAndHash(fakePassword, encodedHash)
	if err != nil {
		t.Errorf("Unexpected error comparing passwords: %s\n", err.Error())
	}
	if match {
		t.Errorf("Comparing passwords: (%s, %s) should have resulted false, got %t\n",
			password, fakePassword, match)
	}

	match, err = argon.ComparePasswordAndHash(password, encodedHash)
	if err != nil {
		t.Errorf("Unexpected error comparing passwords: %s\n", err.Error())
	}
	if !match {
		t.Errorf("Comparing passwords: (%s, %s) should have resulted true, got %t\n",
			password, password, match)
	}
}
