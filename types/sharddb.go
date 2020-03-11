package types

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/xoreo/meros/models"
)

// ShardDB is the database that holds the locations of each shard of a (larger) file.
type ShardDB struct {
	Header *DatabaseHeader    `json:"header"` // The database header contains some DB metadata
	Shards map[peer.ID]*Shard `json:"shards"` // The shards themselves
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

/*
THESE METHODS WILL GO IN THE NET PACKAGE:
// PopulateShardAddresses populates the addresses within the Shards map with peer addresses on the network.
func (shardDB *ShardDB) PopulateShardAddresses() error

// DistributeShards distributes the shards across the network.
func (shardDB *ShardDB) DistributeShards() error
*/
