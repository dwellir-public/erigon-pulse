package pulse

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/chain"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/log/v3"

	"github.com/ledgerwatch/erigon/core/state"
)

// see https://gitlab.com/pulsechaincom/compressed-allocations/-/tags/Mainnet
//
//go:embed sacrifice_credits_mainnet.bin
var mainnetRawCredits []byte

// see https://gitlab.com/pulsechaincom/compressed-allocations/-/tags/Testnet-V4
//
//go:embed sacrifice_credits_testnet_v4.bin
var testnetV4RawCredits []byte

// Applies the sacrifice credits for the PrimordialPulse fork.
func applySacrificeCredits(state *state.IntraBlockState, pulseChainConfig *chain.PulseChainConfig, chainID *big.Int) {
	rawCredits := mainnetRawCredits
	if chainID.Cmp(TestnetV4ChainID) == 0 {
		rawCredits = testnetV4RawCredits
	}

	if pulseChainConfig != nil && pulseChainConfig.Treasury != nil {
		balance, err := uint256.FromHex(pulseChainConfig.Treasury.Balance)
		if err != nil {
			panic(err)
		}
		log.Info("Applying PrimordialPulse treasury allocation 💸")
		log.Info(fmt.Sprintf("Applying PrimordialPulse treasury allocation address: %s", libcommon.HexToAddress(pulseChainConfig.Treasury.Addr).String()))
		log.Info(fmt.Sprintf("Applying PrimordialPulse treasury allocation amount: %d", balance))
		state.AddBalance(libcommon.HexToAddress(pulseChainConfig.Treasury.Addr), balance)
	}

	log.Info("Applying PrimordialPulse sacrifice credits 💸")
	for ptr := 0; ptr < len(rawCredits); {
		byteCount := int(rawCredits[ptr])
		ptr++

		record := rawCredits[ptr : ptr+byteCount]
		ptr += byteCount

		addr := libcommon.BytesToAddress(record[:20])
		hexBalance := hex.EncodeToString(record[20:])
		credit, err := uint256.FromHex("0x" + strings.TrimLeft(hexBalance, "0"))
		if err != nil {
			log.Info(fmt.Sprintf("Applying PrimordialPulse sacrifice credits amount: %s", hexBalance))
			panic(err)
		}
		state.AddBalance(addr, credit)

	}

	log.Info("Finished applying PrimordialPulse sacrifice credits 🤑")
}
