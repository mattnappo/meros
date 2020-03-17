package database

import (
	"encoding/hex"

	"github.com/xoreo/meros/crypto"
)

// ID represents a hash for the keys in the database.
type ID crypto.Hash

// IDFromString returns an ID given a string
func IDFromString(s string) (ID, error) {
	b, err := hex.DecodeString(s) // Decode from hex into []byte
	if err != nil {
		return ID{}, err
	}

	idHash, err := crypto.NewHash(b) // Create the hash
	return ID(idHash), err           // Return the cast to ID
}

// Bytes converts a given hash to a byte array.
func (id ID) Bytes() []byte {
	hash := crypto.Hash(id)
	return hash.Bytes() // Return byte array value
}

// String returns the hash as a hex string.
func (id ID) String() string {
	b := id.Bytes()
	return hex.EncodeToString(b) // Convert to a hex string
}
