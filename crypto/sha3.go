package crypto

/* sha3.go was adapted from @dowlandaiello */

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// HashLength is the standardized length of a hash.
const HashLength = 32

// Hash represents the streamlined hash type to be used.
type Hash [HashLength]byte

// NewHash constructs a new hash given a hash, API so it returns an error.
func NewHash(b []byte) (Hash, error) {
	var hash Hash // Setup the hash
	bCropped := b // Setup the cropped buffer

	// Check the crop side
	if len(b) > len(hash) {
		bCropped = bCropped[len(bCropped)-HashLength:] // Crop the hash
	}

	// Copy the source
	copy(
		hash[HashLength-len(bCropped):],
		bCropped,
	)

	return hash, nil
}

// newHash constructs a new hash given a hash, returns no error
func newHash(b []byte) Hash {
	var hash Hash // Setup the hash
	bCropped := b // Setup the cropped buffer

	// Check the crop side
	if len(b) > len(hash) {
		bCropped = bCropped[len(bCropped)-HashLength:] // Crop the hash
	}

	// Copy the source
	copy(
		hash[HashLength-len(bCropped):],
		bCropped,
	)

	return hash
}

// Sha3 hashes a []byte using sha3.
func Sha3(b []byte) Hash {
	hash := sha3.New256()
	hash.Write(b)
	return newHash(hash.Sum(nil))
}

// Sha3String hashes a given message via sha3 and encodes the hashed message to a hex string.
func Sha3String(b []byte) string {
	b = Sha3(b).Bytes()
	return hex.EncodeToString(b) // Convert to a hex string
}

// HashFromString returns a Hash type given a hex string.
func HashFromString(s string) (Hash, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return Hash{}, err
	}
	return newHash(b), nil
}

// IsNil checks if a given hash is nil.
func (hash Hash) IsNil() bool {
	nilBytes := 0 // Init nil bytes buffer

	// Iterate through the hash, checking for nil bytes
	for _, byteVal := range hash[:] {
		if byteVal == 0 {
			nilBytes++
		}
	}

	return nilBytes == HashLength
}

// Bytes converts a given hash to a byte array.
func (hash Hash) Bytes() []byte {
	return hash[:] // Return byte array value
}

// String returns the hash as a hex string.
func (hash Hash) String() string {
	b := hash.Bytes()
	return hex.EncodeToString(b) // Convert to a hex string
}
