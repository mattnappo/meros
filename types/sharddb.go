package types

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/xoreo/meros/crypto"
)

// errNilDBLabel is returned when a nil label is given.
var errNilDBLabel = errors.New("label for creating a shard database header must not be nil")

// shardDB is the database that holds the locations of each shard of a (larger) file.
type shardDB struct {
	header DatabaseHeader        // Database header
	shards map[shardID]shardData // Shard data map

	hash crypto.Hash // Hash of the entire database
}

// generateShardDB constructs a new shard database.
func generateShardDB(shards []Shard, nodeIDs []NodeID) (*shardDB, error) {
	// Construct the map
	shardMap := make(map[peer.ID]*Shard)

	// Generate and add shard data to the map
	for shard, i := range shards {
		id, data := generateShardEntry(shard)
	}

	// Construct the database
	sharddb := &shardDB{
		NewDatabaseHeader(""), // Generate and set the header
		shardMap,              // Set the shardMap
	}

	return sharddb, nil
}
