package contracts

import (
	contractutils "github.com/silcprotocol/silcchain/contracts/utils"
	evmtypes "github.com/silcprotocol/silcchain/x/vm/types"
)

func LoadInterchainSenderContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("InterchainSender.json")
}
