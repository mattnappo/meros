package common

const (
	// HashLength is the standardized length of a hash
	HashLength = 32
)

// Hash represents the streamlined hash type to be used
type Hash [HashLength]byte

// NewHash constructs a new hash given a hash
func NewHash(b []byte) Hash {
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

// IsNil checks if a given hash is nil
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

// Bytes converts a given hash to a byte array
func (hash Hash) Bytes() []byte {
	return hash[:] // Return byte array value
}
