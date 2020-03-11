package types

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/xoreo/meros/models"
)

// ErrNilDBLabel is returned when a nil label is given.
var ErrNilDBLabel = errors.New("label for creating a shard database header must not be nil")

/*
	TYPE NOTES
	ShardDBs only exist within the File struct
*/

// ShardDB is the database that holds the locations of each shard of a (larger) file.
type ShardDB struct {
	Header *DatabaseHeader    `json:"header"` // The database header contains some DB metadata.
	Shards map[peer.ID]*Shard `json:"shards"` // The shards themselves. Eventually, this will be a BoltDB.
}

// NewShardDB constructs a new database of shards.
func NewShardDB(label string, bytes []byte) (*ShardDB, error) {
	// Check for valid label
	if label == "" {
		return nil, ErrNilDBLabel
	}

	// Check that file is not nil
	if len(bytes) == 0 {
		return nil, ErrNilFileSize
	}

	newHeader, err := NewDatabaseHeader(label) // Construct the database header
	if err != nil {
		return nil, err
	}

	// Generate the actual shards
	_, err = GenerateShards(bytes, models.ShardCount)
	if err != nil {
		return nil, err
	}

	shardMap := make(map[peer.ID]*Shard)

	// Find online peers, put node addresses as keys in map and shards as values

	// Construct the shard itself
	newShardDB := &ShardDB{
		Header: newHeader,
		Shards: shardMap,
	}

	return newShardDB, nil
}

// PopulateShardAddresses populates the addresses within the Shards map with peer addresses on the network.
func (shardDB *ShardDB) PopulateShardAddresses() error {

	return nil
}

// DistributeShards distributes the shards across the network.
func (shardDB *ShardDB) DistributeShards() error {

	return nil
}