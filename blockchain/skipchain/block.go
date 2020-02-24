package skipchain

import (
	"crypto/sha256"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
)

// Block is the representation of the data structures that will be linked
// together.
type Block struct {
	Index     uint64
	Roster    blockchain.Roster
	Signature cosi.Signature
	Data      proto.Message
}

func (b Block) hash() ([]byte, error) {
	h := sha256.New()

	buffer, err := proto.Marshal(b.Data)
	if err != nil {
		return nil, err
	}

	_, err = h.Write(buffer)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Pack returns a network message.
func (b Block) Pack() proto.Message {
	payload, _ := ptypes.MarshalAny(b.Data)
	metadata, _ := ptypes.MarshalAny(&BlockMetaData{
		Signature: b.Signature,
	})

	return &blockchain.VerifiableBlock{
		Block: &blockchain.Block{
			Index:    b.Index,
			Payload:  payload,
			Metadata: metadata,
		},
	}
}

type blockFactory struct {
	verifier cosi.Verifier
}

func (f blockFactory) Create(src *blockchain.VerifiableBlock, pubkeys []kyber.Point) (interface{}, error) {
	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(src.Block.GetPayload(), &da)
	if err != nil {
		return Block{}, err
	}

	var metadata BlockMetaData
	err = ptypes.UnmarshalAny(src.Block.GetMetadata(), &metadata)
	if err != nil {
		return Block{}, err
	}

	block := Block{
		Index:     src.Block.GetIndex(),
		Data:      da.Message,
		Signature: metadata.GetSignature(),
	}

	hash, err := block.hash()
	if err != nil {
		return block, err
	}

	err = f.verifier(pubkeys, hash, block.Signature)
	if err != nil {
		return block, err
	}

	return block, nil
}
