package types

import "testing"

func TestGenerateShardEntry(t *testing.T) {
	shard, err := NewShard([]byte("I am an awesome test file"))
	if err != nil {
		t.Fatal(err)
	}

	node, err := NewNodeID("0.0.0.0", 8000)
	if err != nil {
		t.Fatal(err)
	}

	shardid, shardData := generateShardEntry(*shard, 0, *node)
	t.Logf("\nshardID: %x\nshardData: %v\n", shardid, shardData)
}
