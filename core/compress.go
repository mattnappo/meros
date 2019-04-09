package core

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

// CompressBytes compresses a given []byte using compress/zlib
func CompressBytes(rawBytes []byte) ([]byte, error) {
	var buffer bytes.Buffer           // Init the buffer
	writer := zlib.NewWriter(&buffer) // Create the writer

	_, err := writer.Write(rawBytes) // Compress the bytes
	if err != nil {
		return nil, err
	}

	writer.Close() // Close the writer

	return buffer.Bytes(), nil // Return the bytes
}

// DecompressBytes decompresses bytes given a []byte containing previously compressed bytes
func DecompressBytes(rawBytes []byte) ([]byte, error) {
	var buffer bytes.Buffer          // Init the buffer
	_, err := buffer.Write(rawBytes) // Write the raw to the buffer
	if err != nil {
		return nil, err
	}

	reader, err := zlib.NewReader(&buffer) // Init the reader
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(os.Stdout, reader) // Decompress
	if err != nil {
		return nil, err
	}

	err = reader.Close() // Close the reader
	if err != nil {
		return nil, err
	}

	fmt.Printf("inner: %x\n", reader)

	return buffer.Bytes(), nil
}
