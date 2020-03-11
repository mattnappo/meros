package filedb

import (
	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/common"
)

// FileDB implements the main file database that holds the locations
// for all the files on the network.
type FileDB struct {
	DB *bolt.DB // BoltDB instance
}

// Open opens the database for reading and writing.
func Open() (*FileDB, error) {
	err := common.CreateDirIfDoesNotExist(models.DBPath)
}
