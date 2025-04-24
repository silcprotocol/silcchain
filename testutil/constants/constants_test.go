package constants_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	chainconfig "github.com/silcprotocol/silcchain/cmd/silcd/config"
	"github.com/silcprotocol/silcchain/silcd"
	"github.com/silcprotocol/silcchain/testutil/constants"
)

func TestRequireSameTestDenom(t *testing.T) {
	require.Equal(t,
		constants.ExampleAttoDenom,
		silcd.ExampleChainDenom,
		"test denoms should be the same across the repo",
	)
}

func TestRequireSameTestBech32Prefix(t *testing.T) {
	require.Equal(t,
		constants.ExampleBech32Prefix,
		chainconfig.Bech32Prefix,
		"bech32 prefixes should be the same across the repo",
	)
}

func TestRequireSameWEVMOSMainnet(t *testing.T) {
	require.Equal(t,
		constants.WEVMOSContractMainnet,
		silcd.WEVMOSContractMainnet,
		"wevmos contract addresses should be the same across the repo",
	)
}
