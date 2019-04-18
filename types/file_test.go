package types

import "testing"

func TestNewFile(t *testing.T) {
	filename := "testFilename"
	shardCount := 19
	var size uint32 = 10000 // 10 KB

	file, err := NewFile(filename, shardCount, size)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(file.String())

}
