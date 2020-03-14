package core

import "errors"

// ErrCannotSplitBytes is returned when a byte slice cannot be split evenly
// with the given size array.
var ErrCannotSplitBytes = errors.New(
	"byte array can not be split given size vector",
)

// calculateSizeSum calculates the sum of all of the uint32s in a []uint32.
func calculateSizeSum(sizes []uint32) uint32 {
	var sum uint32 // Init the sum

	// Add all of the slice's elements
	for _, size := range sizes {
		sum += size
	}

	return sum // Return the sum
}

// SplitBytes splits a []byte n times.
func SplitBytes(bytes []byte, sizes []uint32) ([][]byte, error) {
	// Check that the bytes can be split given the size vector
	if uint32(len(bytes)) != calculateSizeSum(sizes) {
		return nil, ErrCannotSplitBytes
	}

	var splitBytes [][]byte // Init the master slice
	currentBytePos := 0     // Init the byte position

	// For each size (shard)
	for _, currentSize := range sizes {
		var tempBytes []byte // Init shard[i]'s byte slice

		// For each byte that needs to be added
		for i := 0; i < int(currentSize); i++ {
			tempBytes = append(tempBytes, bytes[currentBytePos]) // Add the byte

			currentBytePos++ // Move the "byte cursor"
		}

		// Append the shard's bytes to the master slice
		splitBytes = append(splitBytes, tempBytes)
	}

	return splitBytes, nil
}
