package common

import (
	"testing"
)

func TestNewHash(t *testing.T) {
	byteHash := []byte("test")
	newHash := NewHash(byteHash)

	t.Log(newHash)
}
