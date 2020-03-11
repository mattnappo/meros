package types

import "testing"

func TestNewDatabaseHeader(t *testing.T) {
	label := "test label"
	header := NewDatabaseHeader(label)
	t.Log(header.String())
}
