package core

import (
	"bytes"
	"compress/flate"
	"io"
)

// CompressBytes returns a compressed byte slice.
func CompressBytes(src []byte) []byte {
	compressedData := new(bytes.Buffer)
	compress(src, compressedData, -2)
	return compressedData.Bytes()
}

// DecompressBytes returns a decompressed byte slice.
func DecompressBytes(src []byte) []byte {
	compressedData := bytes.NewBuffer(src)
	deCompressedData := new(bytes.Buffer)
	decompress(compressedData, deCompressedData)
	return deCompressedData.Bytes()
}

// compress uses flate to compress a byte slice to a corresponding level.
func compress(src []byte, dest io.Writer, level int) {
	compressor, _ := flate.NewWriter(dest, level)
	compressor.Write(src)
	compressor.Close()
}

// compress uses flate to decompress an io.Reader.
func decompress(src io.Reader, dest io.Writer) {
	decompressor := flate.NewReader(src)
	io.Copy(dest, decompressor)
	decompressor.Close()
}
