package filedb

import (
	"github.com/xoreo/meros/types"
)

// generateFileEntry generates a fileID-file pair for the fileDB.
func generateFileEntry(file types.File) ([]byte, []byte) {
	return file.Hash.Bytes(), file.Bytes()
}
