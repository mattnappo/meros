package types

import (
	"encoding/json"
	"time"

	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/crypto"
)

// DatabaseHeader is the identifier for a database
type DatabaseHeader struct {
	Label   string      `json:"label"`   // A database label
	Created string      `json:"created"` // The time that the database was created
	Hash    common.Hash `json:"hash"`    // The hash of the database header
}

// NewDatabaseHeader creates a new database header
func NewDatabaseHeader(label string) {
	newDatabaseHeader := &DatabaseHeader{
		Label:   label,
		Created: time.Now().String(),
	}

	(*newDatabaseHeader).Hash = crypto.Sha3(newDatabaseHeader.Bytes())
}

/* ----- BEGIN HELPER FUNCTIONS ----- */

// Bytes converts the database header to bytes
func (databaseHeader *DatabaseHeader) Bytes() []byte {
	json, _ := json.MarshalIndent(*databaseHeader, "", "  ")
	return json
}

/* ----- END HELPER FUNCTIONS ----- */
