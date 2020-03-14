package types

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/xoreo/meros/core"
	"github.com/xoreo/meros/crypto"
	"github.com/xoreo/meros/models"
)

var (
	// ErrNilFilename is returned when the fileame to construct a new file is nil.
	ErrNilFilename = errors.New("filename to construct file must not be nil")

	// ErrNilShardCount is returned when the shard counnt to cosntruct a new file is nil.
	ErrNilShardCount = errors.New("shard count to cosntruct file must not be nil")

	// ErrNilFileSize is returned when the file size to construct a new file is nil.
	ErrNilFileSize = errors.New("file size to construct file must not be nil")
)

// File contains the (important) metadata of a file stored in a database.
type File struct {
	Filename   string      `json:"filename"`    // The file's filename
	ShardCount int         `json:"shard_count"` // The number of shards hosting the file
	Size       uint32      `json:"size"`        // Total size of the file
	ShardDB    *shardDB    `json:"shard_db"`    // Pointer to this file's shardDb
	Hash       crypto.Hash `json:"hash"`        // The hash of the file
}

// NewFile constructs a new file from a file in memory.
func NewFile(filename string) (*File, error) {
	// Check that the filename is not nil
	if filename == "" {
		return nil, ErrNilFilename
	}

	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get the file length and bytes
	fileStat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := uint32(fileStat.Size())
	if size == 0 {
		return nil, ErrNilFileSize
	}

	// Read from the file
	bytes := make([]byte, size)
	_, err = f.Read(bytes)
	if err != nil {
		return nil, err
	}

	// Compress the data
	bytes = core.CompressBytes(bytes)
	// Encrypt the data as well (implement later)

	shardCount := models.ShardCount

	// Create a new file pointer
	file := &File{
		Filename:   filename,   // The filename
		ShardCount: shardCount, // The total amount of shards hostinng the file
		Size:       size,       // The total size of the file
		ShardDB:    nil,        // nil for now
	}

	// Compute the hash of the file
	(*file).Hash = crypto.Sha3(file.Bytes())
	return file, nil
}

/* ----- BEGIN HELPER FUNCTIONS ----- */

// Bytes converts the database header to bytes.
func (file *File) Bytes() []byte {
	json, _ := json.MarshalIndent(*file, "", "  ")
	return json
}

// String converts the database to a string.
func (file *File) String() string {
	json, _ := json.MarshalIndent(*file, "", "  ")
	return string(json)
}

/* ----- END HELPER FUNCTIONS ----- */
