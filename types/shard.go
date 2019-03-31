package types

import (
	"errors"

	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/crypto"
)

var (
	// ErrNilBytes is returned when a shard is constructed when given nil bytes
	ErrNilBytes = errors.New("bytes to construct new shard must not be nil")
)

// Shard is a struct that holds a piece of data that is
// a part of another, bigger piece of data
type Shard struct {
	Size  uint32      `json:"size"`  // The size of the shard
	Bytes []byte      `json:"bytes"` // The actual data of the shard
	Hash  common.Hash `json:"hash"`  // The hash of the shard
}

// NewShard attempts to construct a new shard
func NewShard(bytes []byte) (*Shard, error) {
	if bytes == nil {
		return nil, ErrNilBytes
	}

	newShard := &Shard{
		Size:  uint32(len(bytes)),
		Bytes: bytes,
		Hash:  crypto.Sha3(bytes),
	}

	return newShard, nil
}
