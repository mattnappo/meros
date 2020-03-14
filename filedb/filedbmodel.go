package filedb

import (
	"github.com/xoreo/meros/crypto"
	"github.com/xoreo/meros/types"
)

// fileID is the model for keys in the file database.
type fileID crypto.Hash

// generateFileEntry generates a fileID-file pair for the fileDB.
func generateFileEntry(file types.File) (fileID, types.File) {
	return fileID(file.Hash), file
}
