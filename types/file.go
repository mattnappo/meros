package types

// File containsn the metadata and location of the shards in order
// to reconstruct the file
type File struct {
	Filename   string   `json:"filename"`   // The file's filename
	ShardCount int      `json:"shardCount"` // The number of shards hosting the file
	Size       uint32   `json:"size"`       // The total size of the file
	ShardDB    *ShardDB `json:"shardDB"`    // The pointer to the ShardDB, the place where the locations of the shards are
}
