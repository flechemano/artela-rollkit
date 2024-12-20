package api

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	rpctypes "github.com/artela-network/artela-rollkit/ethereum/rpc/types"
)

// web3Api offers network related RPC methods.
type Web3API struct {
	b rpctypes.Web3Backend
}

// NewWeb3API creates a new web3 DebugAPI instance.
func NewWeb3API(b rpctypes.Web3Backend) rpctypes.Web3Backend {
	return &Web3API{b}
}

// ClientVersion returns the node name.
func (api *Web3API) ClientVersion() string {
	return api.b.ClientVersion()
}

// Sha3 applies the ethereum sha3 implementation on the input.
// It assumes the input is hex encoded.
func (*Web3API) Sha3(input hexutil.Bytes) hexutil.Bytes {
	return crypto.Keccak256(input)
}
