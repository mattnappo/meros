package core

import "errors"

// ErrCannotSplitBytes is returned when a byte slice can not be split evenly with the given size array
var ErrCannotSplitBytes = errors.New("byte array can not be split given size vector")

// calculateSizeSum calculates the sum of all of the uint32s in a []uint32
func calculateSizeSum(sizes []uint32) uint32 {
	var sum uint32 // Init the sum

	// Add all of the slice's elements
	for _, size := range sizes {
		sum += size
	}

	return sum // Return the sum
}

// SplitBytes splits a []byte n times
func SplitBytes(bytes []byte, sizes []uint32) ([][]byte, error) {
	// Check that the bytes can be split given the size vector
	if uint32(len(bytes)) != calculateSizeSum(sizes) {
		return nil, ErrCannotSplitBytes
	}

	var splitBytes [][]byte // Init the master slice
	currentBytePos := 0     // Init the byte position

	// currentByte := bytes[currentBytePos]

	for _, currentSize := range sizes {
		var tempBytes []byte // Init shard[i]'s byte slice

		for i := 0; i < int(currentSize); i++ {
			tempBytes = append(tempBytes, bytes[currentBytePos])
			currentBytePos++
		}
		splitBytes = append(splitBytes, tempBytes)

		// at the very end
		// currentBytePos = int(currentSize)
	}
	// for i := 0; i < len(sizes); i++ { // For each shard (size)
	// 	var tempBytes []byte // Init shard[i]'s byte slice

	// 	for j := currentBytePos; j < int(sizes[i]); j++ { // For each byte that the shard should have
	// 		tempBytes = append(tempBytes, bytes[currentBytePos])
	// 		currentBytePos++
	// 	}
	// 	splitBytes = append(splitBytes, tempBytes) // Append to the master slice
	// }
	return splitBytes, nil
}
