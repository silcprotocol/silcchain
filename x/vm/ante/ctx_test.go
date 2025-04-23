package ante_test

import (
	"github.com/silcprotocol/silcchain/testutil/integration/os/network"
	evmante "github.com/silcprotocol/silcchain/x/vm/ante"

	storetypes "cosmossdk.io/store/types"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
