package database

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/crypto"
	"github.com/xoreo/meros/models"
	"github.com/xoreo/meros/types"
)

// DBType represents the type of database of the instance, either a node's
// shard database or the main file database.
type DBType int

const (
	// FILEDB the marker for a file database.
	FILEDB DBType = iota

	// NSHARDDB is the marker for the node's shard database.
	NSHARDDB DBType = iota

	// fileDBBucket is the bucket used to store files within a filedb.
	fileDBBucket = "files"

	// nshardDBBucket is the bucket used to store shards within a shard database.
	nshardDBBucket = "shards"
)

// Database implements a general database that holds various data within meros.
type Database struct {
	Header types.DatabaseHeader `json:"header"` // Database header info
	Name   string               `json:"name"`   // The name of the db
	DBType DBType // The type of database

	DB     *bolt.DB             // BoltDB instance
	bucket []byte // The bucket for the database
	open bool // Status of the DB
}

// Open opens the database for reading and writing. Creates a new DB if one
// with that name does not already exist.
func Open(dbName string, dbType DBType) (*Database, error) {
	// Make sure path exists
	err := common.CreateDirIfDoesNotExist(path.Join(models.DataPath, dbName))
	if err != nil {
		return nil, err
	}

	var database *Database // The database to return

	// Prepare to serialize the database struct
	databasePath := path.Join(models.DBPath, dbName, "db.json")
	if _, err := os.Stat(databasePath); err != nil { // If DB name does not exist
		// Create the database struct
		database = &Database{
			Header: types.NewDatabaseHeader(dbName), // Generate and set the header

			Name: dbName, // Set the name
		}

		err = database.serialize(databasePath) // Write the database struct to disk
		if err != nil {
			return nil, err
		}

	} else {
		// If the db does exist, read from it and return it
		database, err = deserialize(databasePath)
		if err != nil {
			return nil, err
		}
	}

	// Prepare to open the bolt database
	boltdbPath := path.Join(models.DBPath, dbName, "bolt.db")
	db, err := bolt.Open(boltdbPath, 0600, &bolt.Options{ // Open the DB
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	database.DB = db // Set the DB
	database.DBType = dbType // Set the type of database

	// Bucket handler 
	switch dbType {
	case FILEDB: // If FileDB type, set to the corresponding bucket
		database.bucket = fileDBBucket
	case NSHARDDB: // If NodeShardDB type, set to the corresponding bucket
		database.bucket = nshardDBBucket
	}

	err = database.makeBuckets() // Make the buckets in the database
	if err != nil {
		return nil, err
	}

	database.open = true // Set the status to open

	return database, nil
}

// Close closes the database.
func (db *Database) Close() error {
	err := db.DB.Close() // Close the DB
	if err != nil {
		return err
	}

	db.open = false // Set DB status
	return nil
}

// makeBuckets constructs the buckets in the database.
func (db *Database) makeBuckets() error {
	// Create all buckets in the database
	for _, bucket := db.Buckets {
		err := db.DB.Update(func(tx *bolt.Tx) error { // Open tx for bucket creation
			_, err := tx.CreateBucketIfNotExists(db.bucket) // Create bucket
			return err                                        // Handle err
		})
		if err != nil { // Check the err
			return err
		}
	}
}

// String marshals the DB as a string.
func (db *Database) String() string {
	json, _ := json.MarshalIndent(*db, "", "  ")
	return string(json)
}

// serialize will serialize the database and write it to disk.
func (db *Database) serialize(filepath string) error {
	json, _ := json.MarshalIndent(*db, "", "  ")
	err := ioutil.WriteFile(filepath, json, 0600)
	return err
}

// deserialize will deserialize the database from the disk
func deserialize(filepath string) (*Database, error) {
	data, err := ioutil.ReadFile(filepath) // Read the database from disk
	if err != nil {
		return nil, err
	}

	buffer := &Database{} // Initialize the database buffer

	// Unmarshal and write into the buffer
	err = json.Unmarshal(data, buffer)

	return buffer, err
}

