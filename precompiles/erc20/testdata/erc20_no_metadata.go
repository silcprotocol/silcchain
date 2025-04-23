package testdata

import (
	contractutils "github.com/silcprotocol/silcchain/contracts/utils"
	evmtypes "github.com/silcprotocol/silcchain/x/vm/types"
)

func LoadERC20NoMetadataContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20NoMetadata.json")
}
