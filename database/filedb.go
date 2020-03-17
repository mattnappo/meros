package database

// filesBucket is the string representation of the bucket for files.
const filesBucket = "files"

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

// PutFile adds a new file to the database.
func (filedb *Database) PutFile(file types.File) (FileID, error) {
	if filedb.open == false { // Make sure the DB is open
		var t FileID // Temporary nil file ID
		return t, errors.New("filedb is closed")
	}

	// Extract the data for the database
	fileid, filedata := generateFileEntry(file)

	// Write the file to the bucket
	err := filedb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(filesBucket) // Fetch the bucket

		// Put necessary data into the bucket
		return b.Put(fileid.Bytes(), filedata)
	})

	return fileid, err
}

// GetFile gets a file from the database.
func (filedb *Database) GetFile(fileid FileID) (*types.File, error) {
	if filedb.open == false { // Make sure the DB is open
		return nil, errors.New("filedb is closed")
	}

	// Initialize file buffer
	var fileBuffer []byte

	// Read from the database
	err := filedb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(filesBucket)       // Fetch the bucket
		readfile := b.Get(fileid.Bytes()) // Read the file
		if readfile == nil {              // Check the file not nil
			return errors.New(
				"file '" + fileid.String() + "' not found in filedb '" + filedb.Name + "'",
			) // Return err if nil
		}

		fileBuffer = make([]byte, len(readfile)) // Init the buffer size
		copy(fileBuffer, readfile)               // Copy the file to the buffer
		return nil
	})

	// Construct file from bytes and return
	file, err := types.FileFromBytes(fileBuffer)
	return file, err
}
