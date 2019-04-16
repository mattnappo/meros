package types

// DatabaseID is the identifier for a database
type DatabaseID struct {
	Label   string `json:"label"`
	Created []byte `json:"created"`
	Hash    Hash   `json:"hash"`
}
