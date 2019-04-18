package types

import (
	"encoding/json"
	"errors"
)

var (
	// ErrNilFilename is returned when the fileame to construct a new file is nil
	ErrNilFilename = errors.New("filename to construct file must not be nil")

	// ErrNilShardCount is returned when the shard counnt to cosntruct a new file is nil
	ErrNilShardCount = errors.New("shard count to cosntruct file must not be nil")

	// ErrNilFileSize is returned when the file size to construct a new file is nil
	ErrNilFileSize = errors.New("file size to construct file must not be nil")
)

// File containsn the metadata and location of the shards in order
// to reconstruct the file
type File struct {
	Filename   string   `json:"filename"`   // The file's filename
	ShardCount int      `json:"shardCount"` // The number of shards hosting the file
	Size       uint32   `json:"size"`       // The total size of the file
	ShardDB    *ShardDB `json:"shardDB"`    // The pointer to the ShardDB, the place where the locations of the shards are stored
}

// NewFile constructs a new file
func NewFile(filename string, shardCount int, size uint32) (*File, error) {
	// Check that the filename is not nil
	if filename == "" {
		return nil, ErrNilFilename
	}

	// Check that the shard count is not nil
	if shardCount == 0 {
		return nil, ErrNilShardCount
	}

	// Check that the file size is not nil
	if size == 0 {
		return nil, ErrNilFileSize
	}

	// Create a new file pointer
	file := &File{
		Filename:   filename,   // The filename
		ShardCount: shardCount, // The total amount of shards hostinng the file
		Size:       size,
		ShardDB:    nil,
	}
	return file, nil
}

/* ----- BEGIN HELPER FUNCTIONS ----- */

// Bytes converts the database header to bytes
func (file *File) Bytes() []byte {
	json, _ := json.MarshalIndent(*file, "", "  ")
	return json
}

// String converts the database to a string
func (file *File) String() string {
	json, _ := json.MarshalIndent(*file, "", "  ")
	return string(json)
}

/* ----- END HELPER FUNCTIONS ----- */
