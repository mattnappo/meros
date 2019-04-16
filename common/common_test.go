package common

import "testing"

func TestCreateDirIfDoesNotExist(t *testing.T) {
	dir := "test/"
	err := CreateDirIfDoesNotExist(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("created dir '%s'\n", dir)
}
