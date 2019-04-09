package core

import (
	"testing"
)

func TestCompressBytes(t *testing.T) {
	// var b bytes.Buffer
	// w := zlib.NewWriter(&b)
	// w.Write([]byte("hello, world\n")) // Write the compressed bytes to the writer
	// w.Close()

	// t.Logf("%x\n\n", b.String())

	raw := []byte("hello world")
	t.Logf("raw: %x\n", raw)

	compressed, err := CompressBytes(raw)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("compressed: %x\n", compressed)

	decompressed, err := DecompressBytes(compressed)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decompressed: %x\n", decompressed)

}
