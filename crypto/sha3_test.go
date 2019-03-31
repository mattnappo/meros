package crypto

import "testing"

func TestSha3(t *testing.T) {
	payload := []byte("payload")
	if Sha3(payload).isNil() {
		t.Fatal("hash should not be nil")
	}
}

func TestSha3String(t *testing.T) {
	payload := []byte("payload")
	if Sha3String(payload) == "" {
		t.Fatal("hash should not be nil")
	}
}
