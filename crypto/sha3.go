package crypto

import (
	"encoding/hex"

	"github.com/xoreo/meros/common"
	"golang.org/x/crypto/sha3"
)

// Sha3 hashes a []byte using sha3.
func Sha3(b []byte) common.Hash {
	hash := sha3.New256()
	hash.Write(b)
	return common.NewHash(hash.Sum(nil))
}

// Sha3String hashes a given message via sha3 and encodes the hashed message to a hex string.
func Sha3String(b []byte) string {
	b = Sha3(b).Bytes()
	return hex.EncodeToString(b) // Convert to a hex string
}
