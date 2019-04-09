package core

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

// CompressBytes compresses a given []byte using compress/zlib
func CompressBytes(rawBytes []byte) error {
	var buffer *bytes.Buffer // Init the buffer
	// buffer.Write(rawBytes)
	writer := zlib.NewWriter(buffer) // Create the writer

	_, err := writer.Write(rawBytes) // Compress the bytes
	if err != nil {
		return err
	}

	writer.Close() // Close the writer

	// Return the bytes later on
	return nil
}

// DecompressBytes decompresses bytes given a []byte containing previously compressed bytes
func DecompressBytes(rawBytes []byte) {
	// Init the buffer
	var buffer *bytes.Buffer
	buffer.Write(rawBytes)

	reader, err := zlib.NewReader(buffer) // Init the reader
	io.Copy(os.Stdout, reader)            // Decompress

	reader.Close() // Close the reader
}
