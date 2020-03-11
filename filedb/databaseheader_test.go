package filedb

import "testing"

func TestNewDatabaseHeader(t *testing.T) {
	label := "test label"
	header, err := NewDatabaseHeader(label)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(header.String())

}
