package filedb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/models"
	"github.com/xoreo/meros/types"
)

// FileDB implements the main file database that holds the locations
// for all the files on the network.
type FileDB struct {
	Header types.DatabaseHeader `json:"header"` // Database header info
	Name   string               `json:"name"`   // The name of the file db.
	DB     *bolt.DB             // BoltDB instance

	Open bool // Status of the DB
}

// Open opens the database for reading and writing. Creates a new DB if one
// with that name does not already exist.
func Open(dbName string) (*FileDB, error) {
	// Make sure path exists
	err := common.CreateDirIfDoesNotExist(path.Join(models.FileDBPath, dbName))
	if err != nil {
		return nil, err
	}

	/*
		Path will look like this:
			data/
			- file_db
				- filedb1/
					- bolt.db
					- db.json
				- filedb2/
					- bolt.db
					- db.json
			- some_other_db/
			- other_data_info/
	*/

	// Prepare to open the bolt database
	boltdbPath := path.Join(models.FileDBPath, dbName, "bolt.db")
	db, err := bolt.Open(boltdbPath, 0600, &bolt.Options{ // Open the DB
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	var fileDB *FileDB // The fileDB to return

	// Prepare to serizlize the FileDB struct
	filedbPath := path.Join(models.FileDBPath, dbName, "db.json")
	if _, err := os.Stat(filedbPath); err != nil { // If DB name does not exist
		// Generate header info
		header, err := types.NewDatabaseHeader(dbName)

		// Create the fileDB struct
		fileDB = &FileDB{
			Header: header,
			Name:   dbName,
			DB:     db, // Set the DB
			Open:   true,
		}

		fileDB.serialize(filedbPath) // Write the FileDB struct to disk
	}

	return fileDB, nil
}

// Close closes the database.
func (filedb *FileDB) Close() error {
	err := filedb.DB.Close() // Close the DB
	if err != nil {
		return err
	}

	filedb.Open = false // Set DB status
	return nil
}

// PutFile adds a new file to the database.
func (filedb *FileDB) PutFile() {

}

// serialize will serialize the database.
func (filedb *FileDB) serialize(filepath string) error {
	json, _ := json.MarshalIndent(*filedb, "", "  ")
	err := ioutil.WriteFile(filepath, json, 0600)
	return err
}
