package types

import "testing"

func TestNewNodeID(t *testing.T) {
	ip := ""
	port := 3000

	_, err := NewNodeID(ip, port)
	if err != nil {
		t.Fatal(err)
	}

}