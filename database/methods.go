package database

import (
	"errors"

	"github.com/boltdb/bolt"
	"github.com/xoreo/meros/types"
)

// generateEntry generates an ID-file/shard pair for the DB.
func generateEntry(item interface{}) (ID, []byte, error) {
	// Type assertion (disambiguation)
	if t, ok := item.(types.File); ok {
		return ID(t.Hash), t.Bytes(), nil
	} else if t, ok := item.(types.Shard); ok {
		return ID(t.Hash), t.Serialize(), nil
	} else {
		return ID{}, nil, errors.New("invalid type to store in database")
	}
}

// PutItem adds a new item to the database.
func (db *Database) PutItem(item interface{}) (ID, error) {
	var t ID              // Temporary nil item ID
	if db.open == false { // Make sure the DB is open
		return t, errors.New("database is closed")
	}

	// Extract the data for the database
	id, data, err := generateEntry(item)
	if err != nil {
		return ID{}, err
	}

	// Write the item to the bucket
	err = db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.bucket) // Fetch the bucket

		// Put necessary data into the bucket
		return b.Put(id.Bytes(), data)
	})

	return id, err
}

// GetItem gets an item from the database.
func (db *Database) GetItem(id ID) (interface{}, error) {
	if db.open == false { // Make sure the DB is open
		return nil, errors.New("db is closed")
	}

	// Initialize buffer
	var buffer []byte

	// Read from the database
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.bucket)   // Fetch the bucket
		dbRead := b.Get(id.Bytes()) // Read the item from the db
		if dbRead == nil {          // Check the item not nil
			return errors.New(
				"item '" + id.String() + "' not found in db '" + db.Name + "'",
			) // Return err if nil
		}

		buffer = make([]byte, len(dbRead)) // Init the buffer size
		copy(buffer, dbRead)               // Copy the item to the buffer
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Construct corresponding type from bytes and return
	switch db.DBType {
	case FILEDB:
		return types.FileFromBytes(buffer)
	case NSHARDDB:
		return types.ShardFromBytes(buffer)
	}

	// Throw for undefined behavior
	return nil, errors.New("invalid database type was used")
}
