package models

import "path/filepath"

// ShardCount is amount of shards a file should be split into.
const ShardCount int = 3

// MaxFileSize is the maximum size a file on the network can be (in bytes).
const MaxFileSize int = 1000

// DataPath is where all data is held.
const DataPath = "./data"

// FileDBPath is the path of the FileDB.
var FileDBPath = filepath.Join(DataPath, "file_db")
