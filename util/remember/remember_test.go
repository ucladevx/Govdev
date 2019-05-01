package remember

import "testing"

func TestRememberToken(t *testing.T) {
	token, err := RememberToken()
	if err != nil {
		t.Errorf("Error while creating remember token%s\n", err.Error())
	}

	nBytes, err := NBytes(token)

	if nBytes != RememberTokenBytes {
		t.Errorf("Number of bytes should be %d", RememberTokenBytes)
	}
}
