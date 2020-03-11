package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/xoreo/meros/crypto"
)

// ErrNilDBLabel is returned when a nil label is given.
var ErrNilDBLabel = errors.New("label for creating a shard database header must not be nil")

// DatabaseHeader is the identifier for a database.
type DatabaseHeader struct {
	Label   string      `json:"label"`   // A database label
	Created string      `json:"created"` // The time that the database was created
	Hash    crypto.Hash `json:"hash"`    // The hash of the database header
}

// NewDatabaseHeader creates a new database header.
func NewDatabaseHeader(label string) (*DatabaseHeader, error) {
	// Check that the label is not nil
	if label == "" {
		return nil, ErrNilDBLabel
	}

	// Create the header
	newDatabaseHeader := &DatabaseHeader{
		Label:   label,
		Created: time.Now().String(), // The timestamp
	}

	// Compute the header hash and return
	(*newDatabaseHeader).Hash = crypto.Sha3(newDatabaseHeader.Bytes())
	return newDatabaseHeader, nil
}

/* ----- BEGIN HELPER FUNCTIONS ----- */

// Bytes converts the database header to bytes.
func (databaseHeader *DatabaseHeader) Bytes() []byte {
	json, _ := json.MarshalIndent(*databaseHeader, "", "  ")
	return json
}

// String converts the database to a string.
func (databaseHeader *DatabaseHeader) String() string {
	json, _ := json.MarshalIndent(*databaseHeader, "", "  ")
	return string(json)
}

/* ----- END HELPER FUNCTIONS ----- */