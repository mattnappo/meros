package types

import (
	"fmt"
	"net"

	"github.com/xoreo/meros/crypto"
)

// ErrInvalidIP is an error thrown when the ip to construct a NodeID is invalid.
var ErrInvalidIP = errors.New("ip to construct a NodeID is invalid")

// NodeID contains the necessary data for referencing and connecting to a node.
type NodeID struct {
	IP   string      `json:"ip"`   // The node's IP address
	Port int         `json:"port"` // The port on which the node is hosted
	Hash crypto.Hash `json:"hash"` // The hash of the node
}

// NewNodeID constructs a new NodeID.
func NewNodeID(ip string, port int) (*NodeID, error) {
	// Check that the IP address is valid
	addr := net.ParseIP(ip)
	if addr == nil {
		return ErrInvalidIP
	}

	return nil, nil
}
