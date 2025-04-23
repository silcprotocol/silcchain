package contracts

import (
	contractutils "github.com/silcprotocol/silcchain/contracts/utils"
	evmtypes "github.com/silcprotocol/silcchain/x/vm/types"
)

func LoadReverterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("Reverter.json")
}
