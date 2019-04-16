package types

import (
	"testing"

	"github.com/xoreo/meros/crypto"
)

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

func TestFromBytes(t *testing.T) {
	bytes := []byte("test bytes")
	shard, err := NewShard(bytes)
	if err != nil {
		t.Fatal(err)
	}

	shardBytes := shard.Serialize()

	newShard, err := ShardFromBytes(shardBytes)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[newShard] %s\n", newShard.String())

}

func TestValidate(t *testing.T) {
	bytes := []byte("test bytes")
	shard, err := NewShard(bytes)
	if err != nil {
		t.Fatal(err)
	}

	isValid := (*shard).Validate()
	if isValid == false {
		t.Fatal("shard is actually valid")
	}

	badHash := crypto.Sha3([]byte("not a valid hash"))
	(*shard).Hash = badHash

	isValid = (*shard).Validate()
	if isValid == true {
		t.Fatal("shard is actually invalid")
	}
}
