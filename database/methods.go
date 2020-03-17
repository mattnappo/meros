package database

import (
	"encoding/hex"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/crypto"
	"github.com/xoreo/meros/types"
)

// generateEntry generates an ID-file/shard pair for the DB.
func generateEntry(item interface{}) (ID, []byte) {
	return ID(item.Hash), item.Bytes()
}

// PutItem adds a new item to the database.
func (db *Database) PutItem(item interface{}) (ID, error) {
	var t ID              // Temporary nil item ID
	if db.open == false { // Make sure the DB is open
		return t, errors.New("database is closed")
	}

	// Extract the data for the database
	id, data := generateEntry(item)

	// Write the item to the bucket
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.bucket) // Fetch the bucket

		// Put necessary data into the bucket
		return b.Put(id.Bytes(), data)
	})

	return id, err
}

// GetFile gets a file from the database.
func (filedb *Database) GetFile(id ID) (*types.File, error) {
	if filedb.open == false { // Make sure the DB is open
		return nil, errors.New("filedb is closed")
	}

	// Initialize file buffer
	var fileBuffer []byte

	// Read from the database
	err := filedb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(filedb.bucket) // Fetch the bucket
		readfile := b.Get(id.Bytes()) // Read the file
		if readfile == nil {          // Check the file not nil
			return errors.New(
				"file '" + id.String() + "' not found in filedb '" + filedb.Name + "'",
			) // Return err if nil
		}

		fileBuffer = make([]byte, len(readfile)) // Init the buffer size
		copy(fileBuffer, readfile)               // Copy the file to the buffer
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Construct file from bytes and return
	file, err := types.FileFromBytes(fileBuffer)
	return file, err
}
