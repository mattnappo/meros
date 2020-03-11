package types

import (
	"math/rand"
	"testing"

	"github.com/xoreo/meros/crypto"
)

func TestGenerateShardDB(t *testing.T) {
	// Generate shards
	testBytes := []byte("im a neat test file")
	shards, err := GenerateShards(testBytes, 3)
	if err != nil {
		t.Fatal(err)
	}

	// Make len(shards) random nodes
	var nodes []NodeID
	for i := 0; i < len(shards); i++ {
		randomPort := rand.Intn(10000-9000) + 9000
		randomNode, err := NewNodeID("0.0.0.0", randomPort)
		if err != nil {
			t.Fatal(err)
		}

		nodes = append(nodes, *randomNode)
	}

	// Generate shardDB
	sharddb, err := generateShardDB(shards, nodes)
	if err != nil {
		t.Fatal(err)
	}

	// Test the map
	for k, v := range sharddb.shardMap {
		t.Logf("%x: %v\n\n", k, v)
	}

	// Test map access
	h, err := crypto.HashFromString("bafc4c93a862aecc87368f291090fc2fe479eecc3bd15b7efdde01ff92c42592")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\nGET: [%v]\n", sharddb.shardMap[shardID(h)])

	// Analyze shard sizes
	var sizeCounter uint32
	for _, v := range sharddb.shardMap {
		t.Logf("size: %d\n", v.Size)
		sizeCounter = sizeCounter + v.Size
	}
	t.Logf("correct size: %d", len(testBytes))
	t.Logf(" actual size: %d", sizeCounter)
}
