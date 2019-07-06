package types

import (
	"encoding/json"
	"errors"
	"net"

	"github.com/xoreo/meros/crypto"
)

var (
	// ErrInvalidIP is an error thrown when the ip to construct a NodeID is invalid.
	ErrInvalidIP = errors.New("ip to construct a NodeID is invalid")

	// ErrInvalidPort is an error thrown when the port to construct a NodeID is invalid.
	ErrInvalidPort = errors.New("port to construct a NodeID is invalid")
)

// NodeID contains the necessary data for referencing and connecting to a node.
type NodeID struct {
	IP   string       `json:"ip"`   // The node's IP address
	Port int          `json:"port"` // The port on which the node is hosted
	Hash *crypto.Hash `json:"hash"` // The hash of the node
}

// NewNodeID constructs a new NodeID.
func NewNodeID(ip string, port int) (*NodeID, error) {
	// Check that the IP address is valid
	addr := net.ParseIP(ip)
	if addr == nil {
		return nil, ErrInvalidIP
	}

	// Check that the port is valid
	if port == 0 {
		return nil, ErrInvalidPort
	}

	// Create the NodeID
	newNodeID := &NodeID{
		IP:   ip,
		Port: port,
		Hash: nil,
	}

	hash := crypto.Sha3(newNodeID.Bytes())
	newNodeID.Hash = &hash

	return newNodeID, nil
}

/* ----- BEGIN HELPER FUNCTIONS ----- */

// Bytes returns the bytes of a NodeID.
func (nodeID *NodeID) Bytes() []byte {
	json, _ := json.MarshalIndent(*nodeID, "", "  ")
	return json
}

// String converts the NodeID to a string.
func (nodeID *NodeID) String() string {
	json, _ := json.MarshalIndent(*nodeID, "", "  ")
	return string(json)
}

/* ----- END HELPER FUNCTIONS ----- */
