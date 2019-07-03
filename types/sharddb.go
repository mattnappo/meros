package types

import (
	"errors"

	"github.com/boltdb/bolt"
)

var (
	// ErrNilDBLabel is returned when a nil label is given.
	ErrNilDBLabel = errors.New("label for creating a shard database header must not be nil")
)

// ShardDB is the database that holds the locations of each shard of a (larger) file.
type ShardDB struct {
	Header *DatabaseHeader `json:"header"`
	DB     *bolt.DB        `json:"db"`
}

// NewShardDB constructs a new database of shards.
func NewShardDB(label string) (*ShardDB, error) {
	// Check for valid label
	if label == "" {
		return nil, ErrNilDBLabel
	}

	newHeader, err := NewDatabaseHeader(label) // Construct the database header
	if err != nil {
		return nil, err
	}

	// Construct the shard itself
	newShardDB := &ShardDB{
		Header: newHeader,
		DB:     nil,
	}

	return newShardDB, nil
}
