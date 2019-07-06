package types

import "testing"

func TestNewNodeID(t *testing.T) {
	ip := "192.168.8.1"
	port := 3000

	nodeID, err := NewNodeID(ip, port)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(nodeID)
}
