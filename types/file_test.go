package types

import (
	"io/ioutil"
	"testing"

	"github.com/xoreo/meros/common"
)

func TestNewFile(t *testing.T) {
	filename := "./data/testFilename"
	common.CreateDirIfDoesNotExist("./data")

	testData := []byte("this is a test file.\n\nit is just like any other file that would be transferred\nover the meros network.")
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
