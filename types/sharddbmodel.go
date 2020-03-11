package types

import (
	"crypto"
)

// shardID is the key model in the shard map (shard db).
type shardID crypto.Hash

// shardData is the value model in the shard map (shard db).
type shardData struct {
	NodeID NodeID `json:"node_id"` // The ID of the node holding the shard
	Index  int    `json:"index"`   // The position of the shard out of all the shards for this file
	Size   uint32 `json:"size"`    // The size in bytes of the shard
}

func generateShardEntry(shard Shard) (shardData, shardID) {

}
