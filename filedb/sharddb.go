package filedb

import (
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/models"
)

// FileDB implements the main file database that holds the locations
// for all the files on the network.
type FileDB struct {
	DB   *bolt.DB // BoltDB instance
	Open bool     // Status of the DB
}

// Open opens the database for reading and writing.
func Open() (*FileDB, error) {
	err := common.CreateDirIfDoesNotExist(models.FileDBPath) // Make sure path exists
	if err != nil {
		return nil, err
	}

	dbPath := path.Join(models.FileDBPath, "file_db.db") // Prepare the path
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{    // Open the DB
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	fileDB := &FileDB{
		DB:   db, // Set the DB
		Open: true,
	}

	return fileDB, nil
}
