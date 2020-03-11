package types

import (
	"errors"

	"github.com/xoreo/meros/types"
)

// errNilDBLabel is returned when a nil label is given.
var errNilDBLabel = errors.New("label for creating a shard database header must not be nil")

// shardDB is the database that holds the locations of each shard of a (larger) file.
type shardDB struct {
	header DatabaseHeader        // Database header
	shards map[shardID]shardData // Shard data map

	hash Crypto.Hash // Hash of the entire database
}

// generateShardDB constructs a new database of shards.
func generateShardDB(shards []Shard) (*shardDB, error) {
	// Construct the map
	for shard, i := range shards {

	}

	// Construct the database
	sharddb := &shardDB{
		types.NewDatabaseHeader(""), // Generate and set the header

	}

	return sharddb, nil
}
