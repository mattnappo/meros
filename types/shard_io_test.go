package types

import "testing"

func TestWriteShardToMemory(t *testing.T) {
	rawBytes := []byte("shard test bytes")
	shard, err := NewShard(rawBytes)
	if err != nil {
		t.Fatal(err)
	}

	shard.WriteShardToMemory()

}

func TestReadShardFromMemory(t *testing.T) {
	shard, err := ReadShardFromMemory("57c088cc")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[shardFromMemory] %s\n", (*shard).String())
}
