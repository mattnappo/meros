package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/xoreo/meros/common"
)

// ErrInvalidShard is returned when a shard's hash is not valid when loading from memory.
var ErrInvalidShard = errors.New("shard from memory is not valid")

// WriteShardToMemory writes a shard to memory.
func (shard *Shard) WriteShardToMemory() error {
	bytes := shard.Serialize()

	// Create a dir to store the shards
	err := common.CreateDirIfDoesNotExist("data/shards")
	if err != nil {
		return err
	}

	// Create the filename of the hash
	shardHashString := (*shard).Hash.String()[0:8]
	filename := fmt.Sprintf("data/shards/shard_%s.json", shardHashString)

	err = ioutil.WriteFile(filepath.FromSlash(filename), bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ReadShardFromMemory reads a shard from memory.
func ReadShardFromMemory(hash string) (*Shard, error) {
	// Read the file from memory
	data, err := ioutil.ReadFile(fmt.Sprintf("data/shards/shard_%s.json", hash))
	if err != nil {
		return nil, err
	}

	buffer := &Shard{} // Init a shard buffer

	// Read into the buffer
	err = json.Unmarshal(data, buffer)
	if err != nil {
		return nil, err
	}

	if (*buffer).Validate() == false {
		return nil, ErrInvalidShard
	}

	return buffer, nil // Return the shard pointer
}
