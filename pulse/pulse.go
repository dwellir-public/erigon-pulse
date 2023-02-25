// Package pulse implements the PulseChain fork.
package pulse

import (
	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/core/state"
)

// PrimordialPulseFork applies the PrimordialPulse fork changes.
func PrimordialPulseFork(state *state.IntraBlockState, pulseChainConfig *chain.PulseChainConfig) {
	applySacrificeCredits(state, pulseChainConfig)
	replaceDepositContract(state)
}
