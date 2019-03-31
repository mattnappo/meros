package common

import (
	"testing"
)

func TestNewHash(t *testing.T) {
	byteHash := []byte("11cd54753fc9e1d82e39f3b6f9727a3cc4cdf58eec127ccbe056829b1e0a9962")
	newHash := NewHash(byteHash)

	t.Log(newHash)
}

func TestIsNil(t *testing.T) {
	byteHash := []byte("11cd54753fc9e1d82e39f3b6f9727a3cc4cdf58eec127ccbe056829b1e0a9962")
	hash := NewHash(byteHash)

	nilByteHash := []byte("")
	nilHash := NewHash(nilByteHash)

	if hash.IsNil() {
		t.Fatal("hash is not actually nil")
	}

	if nilHash.IsNil() == false {
		t.Fatal("hash is actually nil")
	}
}

func TestBytes(t *testing.T) {
	byteHash := []byte("11cd54753fc9e1d82e39f3b6f9727a3cc4cdf58eec127ccbe056829b1e0a9962")
	hash := NewHash(byteHash)

	t.Log(hash.Bytes())
}

func TestString(t *testing.T) {
	byteHash := []byte("11cd54753fc9e1d82e39f3b6f9727a3cc4cdf58eec127ccbe056829b1e0a9962")
	hash := NewHash(byteHash)

	t.Log(hash.String())
}
