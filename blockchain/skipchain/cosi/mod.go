package cosi

import (
	"context"
	"errors"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/bls"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/onet"
)

var suite = pairing.NewSuiteBn256()

// Verifier is the function used to make sure a signature matches the message
// with a specific list of identities.
type Verifier func(block blockchain.Block, sig []byte) error

// Signature is the response type of a collective signing protocol.
type Signature []byte

// CollectiveSigning is the interface that provides the primitives to sign
// a message by members of a network.
type CollectiveSigning interface {
	Sign(block blockchain.Block) (Signature, error)
	MakeVerifier() Verifier
}

// Validator is the interface that is used to validate a block.
type Validator interface {
	Validate(block blockchain.Block) error
}

// BlsCoSi is an implementation of the collective signing interface by
// using BLS signatures.
type BlsCoSi struct {
	onet onet.Onet
}

// NewBlsCoSi returns a new collective signing instance.
func NewBlsCoSi(o onet.Onet, v Validator) *BlsCoSi {
	identity := o.Identity()
	secretKey := identity.(*key.Pair).Private

	h := func(ctx context.Context, msg onet.Message) (onet.Message, error) {
		switch value := msg.(type) {
		case blockchain.Block:
			err := v.Validate(value)
			if err != nil {
				return nil, err
			}

			buf, err := value.Payload.MarshalBinary()
			if err != nil {
				return nil, err
			}

			return bls.Sign(suite, secretKey, buf)
		}

		return nil, errors.New("unknown request message")
	}

	return &BlsCoSi{
		onet: o.MakeHandler("cosi", h),
	}
}

// Sign returns the collective signature of the block.
func (cosi *BlsCoSi) Sign(block blockchain.Block) (Signature, error) {
	msgs, err := cosi.onet.Collect(block)
	if err != nil {
		return nil, err
	}

	var agg []byte
	ok := true
	var resp onet.Message
	for ok {
		resp, ok = <-msgs
		if ok {
			sig := resp.([]byte)

			if agg == nil {
				agg = sig
			} else {
				agg, err = bls.AggregateSignatures(suite, agg, sig)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return agg, nil
}

// MakeVerifier returns a verifier that can be used to verify signatures
// from this collective signing.
func (cosi *BlsCoSi) MakeVerifier() Verifier {
	return blsVerifier
}

// BlsVerifier verifies that a signature matches the message for the roster public keys.
func blsVerifier(block blockchain.Block, sig []byte) error {
	points := make([]kyber.Point, 0)

	for _, identity := range block.Roster {
		points = append(points, identity.(*key.Pair).Public)
	}

	publicKey := bls.AggregatePublicKeys(suite, points...)

	msg, err := block.Payload.MarshalBinary()
	if err != nil {
		return err
	}

	err = bls.Verify(suite, publicKey, msg, sig)
	if err != nil {
		return err
	}

	return nil
}