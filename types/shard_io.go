package types

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/xoreo/meros/common"
)

// WriteShardToMemory writes a shard to memory
func (shard *Shard) WriteShardToMemory() error {
	bytes := shard.Serialize()

	// Create a dir to store the shards
	err := common.CreateDirIfDoesNotExist("data/shards")
	if err != nil {
		return err
	}

	// Create the filename of the hash
	shardHashString := (*shard).Hash.String()
	filename := fmt.Sprintf("data/shard_%s.json", shardHashString)[0:8]

	err = ioutil.WriteFile(filepath.FromSlash(filename), bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
