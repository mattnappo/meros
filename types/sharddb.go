package types

import "github.com/boltdb/bolt"

// ShardDB is the database that holds the locations of each shard of a (larger) file
type ShardDB struct {
	Header DatabaseHeader `json:"header"`
	DB     *bolt.DB       `json:"db"`
}

// NewShardDB constructs a new database of shards
func NewShardDB() {

}
