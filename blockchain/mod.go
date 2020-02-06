package blockchain

import (
	"context"
	"encoding"

	"go.dedis.ch/phoenix/utils"
)

// Payload is the data structure that can be stored in the chain.
type Payload interface {
	encoding.BinaryMarshaler
}

type Roster []interface{}

// Block is the data structure of the blocks.
type Block struct {
	Index   int64
	Roster  Roster
	Payload Payload
}

// Proof is the interface that provides the primitives to verify that a
// block is valid w.r.t. the genesis block.
type Proof interface {
	// LatestBlock returns the last block of the proof.
	LatestBlock() Block

	// Verify makes sure that the integrity of the block from the genesis block
	// is correct.
	Verify() error
}

// Event is the data structure sent back to observers.
type Event struct {
	Block Block
}

// Blockchain is the interface that provides the primitives to interact with the
// blockchain.
type Blockchain interface {
	// Store stores any representation of a data structure into a new block.
	// The implementation is responsible for any validations required.
	Store(data Payload) error

	// GetProof returns a valid proof of the latest block.
	GetProof() (Proof, error)

	// Watch takes an observer that will be notified for each new block
	// definitely appended to the chain.
	Watch(ctx context.Context, obs utils.Observer)
}