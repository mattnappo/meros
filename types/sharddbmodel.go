package types

import (
	"github.com/xoreo/meros/crypto"
)

// shardID is the key model in the shard map (shard db).
type shardID crypto.Hash

// shardData is the value model in the shard map (shard db).
type shardData struct {
	Index  int    `json:"index"`   // The position of the shard out of all the shards for this file
	NodeID NodeID `json:"node_id"` // The ID of the node holding the shard
	Size   uint32 `json:"size"`    // The size in bytes of the shard
}

func calculateShardID(shard Shard) shardID {

}

// generateShardEntry generates a shardID-shardData pair.
func generateShardEntry(shard Shard, index int, nodeID NodeID) (shardID, shardData) {
	shardData := shardData{
		Index:  index,      // Set the index
		NodeID: nodeID,     // Set the NodeID
		Size:   shard.Size, // Get and set the size
	}
	shardid := calculateShardID(shard) // Calculate the shardID of the shard

	return shardid, shardData // Return the pair
}
