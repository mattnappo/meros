package types

import (
	"io/ioutil"
	"testing"
)

func TestNewFile(t *testing.T) {
	filename := "./data/testFilename"

	testData := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filename, testData, 0644)
	if err != nil {
		t.Fatal(err)
	}

	file, err := NewFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(file.String())

}
