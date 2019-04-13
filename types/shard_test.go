package types

import "testing"

func TestCalculateShardSizes(t *testing.T) {
	rawBytes := []byte("123456789")

	nodes := 5

	sizes, _ := CalculateShardSizes(rawBytes, nodes)
	t.Log(sizes)

}
