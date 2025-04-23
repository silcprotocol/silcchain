package main

import (
	"fmt"
	"os"

	"github.com/silcprotocol/silcchain/cmd/evmd/cmd"
	evmdconfig "github.com/silcprotocol/silcchain/cmd/evmd/config"
	examplechain "github.com/silcprotocol/silcchain/evmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	setupSDKConfig()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "evmd", examplechain.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}

func setupSDKConfig() {
	config := sdk.GetConfig()
	evmdconfig.SetBech32Prefixes(config)
	config.Seal()
}
