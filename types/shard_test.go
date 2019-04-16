package types

import "testing"

func TestNewShard(t *testing.T) {
	bytes := []byte("test bytes")
	newShard, err := NewShard(bytes)
	if err != nil {
		t.Fatal(err)
	}

	t.Log((*newShard).String())
}
func TestCalculateShardSizes(t *testing.T) {
	rawBytes := []byte("123456789")
	nodes := 5

	sizes, _ := CalculateShardSizes(rawBytes, nodes)
	t.Log(sizes)

}

func TestGenerateShards(t *testing.T) {
	bytes := []byte("these are the bytes of a test file that is going to be sharded.")
	nodes := 5

	shards, err := GenerateShards(
		bytes,
		nodes,
	)
	if err != nil {
		t.Fatal(err)
	}

	for i, shard := range shards {
		t.Logf("\n[shard %d] %s\n", i, shard.String())
	}

}
