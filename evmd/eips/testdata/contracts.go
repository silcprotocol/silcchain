package testdata

import (
	contractutils "github.com/silcprotocol/silcchain/contracts/utils"
	evmtypes "github.com/silcprotocol/silcchain/x/vm/types"
)

func LoadCounterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("Counter.json")
}

func LoadCounterFactoryContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("CounterFactory.json")
}
