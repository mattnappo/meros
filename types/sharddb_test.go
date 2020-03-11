package types

import (
	"math/rand"
	"testing"

	"github.com/xoreo/meros/crypto"
)

func TestGenerateShardDB(t *testing.T) {
	shards, err := GenerateShards([]byte("im a neat test file"), 3)
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

	sharddb, err := generateShardDB(shards, nodes)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range sharddb.shardMap {
		t.Logf("%x: %v\n\n", k, v)
	}

	h, err := crypto.HashFromString("bafc4c93a862aecc87368f291090fc2fe479eecc3bd15b7efdde01ff92c42592")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\nGET: [%v]\n", sharddb.shardMap[shardID(h)])

	for _, v := range sharddb.shardMap {
		t.Logf("size: %d\n", v.Size)
	}

}
