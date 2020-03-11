package types

import (
	"errors"

	"github.com/boltdb/bolt"
)

// errNilDBLabel is returned when a nil label is given.
var errNilDBLabel = errors.New("label for creating a shard database header must not be nil")

// shardDB is the database that holds the locations of each shard of a (larger) file.
type shardDB struct {
	DB *bolt.DB // BoltDB instance (shard map)

	Open bool // DB status
}

// newShardDB constructs a new database of shards.
func newShardDB(shards []Shard) (*shardDB, error) {

	// Construct the database
	newShardDB := &shardDB{db, true}

	return newShardDB, nil
}

// populateShardAddresses populates the addresses within the Shards map with peer addresses on the network.
func (shardDB *shardDB) populateShardAddresses() error {

	return nil
}

// distributeShards distributes the shards across the network.
func (shardDB *shardDB) distributeShards() error {

	return nil
}
