package types

import "testing"

func TestNewShard(t *testing.T) {
	bytes := []byte("test bytes")
	newShard, err := NewShard(bytes)
	if err != nil {
		t.Fatal(err)
	}

	t.Log((*newShard).Hash.String())
	t.Log((*newShard).Size)
	t.Log((*newShard).Bytes)
}
func TestCalculateShardSizes(t *testing.T) {
	rawBytes := []byte("123456789")
	nodes := 5

	sizes, _ := CalculateShardSizes(rawBytes, nodes)
	t.Log(sizes)

}

func TestGenerateShards(t *testing.T) {

}
