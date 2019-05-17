package hmac

import "testing"

const testKey = "testKey"
const inputStr = "inputStr"
const wrongStr = "wrongStr"

func TestHash(t *testing.T) {
	h := NewHMAC(testKey)
	hash := h.Hash(inputStr)
	wrongHash := h.Hash(wrongStr)
	secondHash := h.Hash(inputStr)

	if hash == wrongHash {
		t.Errorf("Different strings produced same hash\n")
	}

	if hash != secondHash {
		t.Errorf("Same input string should produce same hash\n")
	}
}
