package filedb

import "testing"

func TestOpenAndClose(t *testing.T) {
	filedb, err := Open("myfiledb")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("filedb: %s\n\n", filedb.String())

	err = filedb.Close()
	if err != nil {
		t.Fatal(err)
	}
}
