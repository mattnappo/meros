package core

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

// CompressBytes compresses a given []byte using compress/zlib
func CompressBytes(rawBytes []byte) (*bytes.Buffer, error) {
	var buffer bytes.Buffer           // Init the buffer
	writer := zlib.NewWriter(&buffer) // Create the writer

	_, err := writer.Write(rawBytes) // Compress the bytes
	if err != nil {
		return nil, err
	}

	writer.Close() // Close the writer

	return &buffer, nil // Return the bytes
}

// DecompressBytes decompresses bytes given a []byte containing previously compressed bytes
func DecompressBytes(rawBytes *bytes.Buffer) ([]byte, error) {
	r, err := zlib.NewReader(rawBytes)
	if err != nil {
		return nil, err
	}
	io.Copy(os.Stdout, r)
	r.Close()
	fmt.Printf("%x\n", r)
	return []byte(""), nil
}
