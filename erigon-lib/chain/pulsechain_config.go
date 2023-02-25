package chain

import (
	"math/big"
)

type PulseChainConfig struct {
	// An optional treasury which will receive allocations during the PrimordialPulseBlock.
	Treasury *PulseChainTreasury `json:"treasury,omitempty"`
}

// String implements the stringer interface, returning the consensus engine details.
func (b *PulseChainConfig) String() string {
	return "PulseChain"
}

// PulseChainTreasury represents an optional treasury for launching PulseChain testnets.
type PulseChainTreasury struct {
	Addr    string `json:"addr"`
	Balance string `json:"balance"`
}

// PulseChainTTDOffset is a trivially small amount of work added to the Ethereum Mainnet TTD
// to allow for un-merging and merging with the PulseChain beacon chain.
var PulseChainTTDOffset = big.NewInt(131_072)
