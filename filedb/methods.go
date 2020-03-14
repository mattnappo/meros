package filedb

import (
	"errors"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/types"
)

// PutFile adds a new file to the database.
func (filedb *FileDB) PutFile(file types.File) error {
	if filedb.open == false { // Make sure the DB is open
		return errors.New("filedb is closed")
	}

	// Write the file to the bucket
	err := filedb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(filesBucket) // Read the bucket

		// Extract the data for the database and put it in the bucket
		return b.Put(generateFileEntry(file))
	})

	return err
}
