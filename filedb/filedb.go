package filedb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/models"
	"github.com/xoreo/meros/types"
)

// filesBucket represents the bucket of files in the database.
var filesBucket = []byte("files")

// FileDB implements the main file database that holds the locations
// for all the files on the network.
type FileDB struct {
	Header types.DatabaseHeader `json:"header"` // Database header info
	Name   string               `json:"name"`   // The name of the file db
	DB     *bolt.DB             // BoltDB instance (file map)

	open bool // Status of the DB
}

// Open opens the database for reading and writing. Creates a new DB if one
// with that name does not already exist.
func Open(dbName string) (*FileDB, error) {
	// Make sure path exists
	err := common.CreateDirIfDoesNotExist(path.Join(models.FileDBPath, dbName))
	if err != nil {
		return nil, err
	}

	var fileDB *FileDB // The fileDB to return

	// Prepare to serizlize the FileDB struct
	filedbPath := path.Join(models.FileDBPath, dbName, "db.json")
	if _, err := os.Stat(filedbPath); err != nil { // If DB name does not exist
		// Create the fileDB struct
		fileDB = &FileDB{
			Header: types.NewDatabaseHeader(dbName), // Generate and set the header

			Name: dbName, // Set the name
		}

		fileDB.serialize(filedbPath) // Write the FileDB struct to disk
	} else {
		// If the db does exist, read from it and return it
		fileDB, err = deserialize(filedbPath)
		if err != nil {
			return nil, err
		}
	}

	// Prepare to open the bolt database
	boltdbPath := path.Join(models.FileDBPath, dbName, "bolt.db")
	db, err := bolt.Open(boltdbPath, 0600, &bolt.Options{ // Open the DB
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	fileDB.makeBuckets() // Make the buckets in the database

	fileDB.DB = db     // Set the DB
	fileDB.open = true // Set the status to open

	return fileDB, nil
}

// Close closes the database.
func (filedb *FileDB) Close() error {
	err := filedb.DB.Close() // Close the DB
	if err != nil {
		return err
	}

	filedb.open = false // Set DB status
	return nil
}

// AddFile addsa a file to the filedb.
func (filedb *FileDB) AddFile(file types.File) error {
	if filedb.open == false { // Make sure the DB is open
		return errors.New("filedb is not open")
	}

	filedb.DB.

}

// makeBuckets constructs the buckets in the file database.
func (filedb *FileDB) makeBuckets() {
	return database.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(filesBucket) // Initialize files bucket
		if err != nil {
			return err
		}
	}
}

// String marshals the DB as a string.
func (filedb *FileDB) String() string {
	json, _ := json.MarshalIndent(*filedb, "", "  ")
	return string(json)
}

// serialize will serialize the database.
func (filedb *FileDB) serialize(filepath string) error {
	json, _ := json.MarshalIndent(*filedb, "", "  ")
	err := ioutil.WriteFile(filepath, json, 0600)
	return err
}

func deserialize(filepath string) (*FileDB, error) {
	data, err := ioutil.ReadFile(filepath) // Read the file from disk
	if err != nil {
		return nil, err
	}

	buffer := &FileDB{} // Initialize a buffer

	// Read(write) into the buffer
	err = json.Unmarshal(data, buffer)

	return buffer, err
}
