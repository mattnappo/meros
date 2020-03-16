package database

// FileID represents a hash for the keys of files in the filedb.
type FileID crypto.Hash

// FileIDFromString returns a FileID given a string
func FileIDFromString(s string) (FileID, error) {
	b, err := hex.DecodeString(s) // Decode from hex into []byte
	if err != nil {
		return FileID{}, err
	}

	fileIDHash, err := crypto.NewHash(b) // Create the hash
	return FileID(fileIDHash), err       // Return the cast to FileID
}

// Bytes converts a given hash to a byte array.
func (fileid FileID) Bytes() []byte {
	hash := crypto.Hash(fileid)
	return hash.Bytes() // Return byte array value
}

// String returns the hash as a hex string.
func (fileid FileID) String() string {
	b := fileid.Bytes()
	return hex.EncodeToString(b) // Convert to a hex string
}

// generateFileEntry generates a fileID-file pair for the fileDB.
func generateFileEntry(file types.File) (FileID, []byte) {
	return FileID(file.Hash), file.Bytes()
}
