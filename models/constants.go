package models

import "path"

// ShardCount is amount of shards a file should be split into.
const ShardCount int = 3

// MaxFileSize is the maximum size a file on the network can be (in bytes).
const MaxFileSize int = 1000

// DataPath is where all data is held.
const DataPath = "./data"

// DBPath represents the path for all databases.
const DBPath = (func() string {
	return path.Join(DataPath, "db")
})

// FileDBPath is the path of the FileDB.
var FileDBPath = path.Join(DataPath, "file_db")

// NodeDBPath is the path of the NodeDB.
var NodeDBPath = path.Join(DataPath, "node_db")
