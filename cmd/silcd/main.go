package main

import (
	"fmt"
	"os"

	"github.com/silcprotocol/silcchain/cmd/silcd/cmd"
	silcdconfig "github.com/silcprotocol/silcchain/cmd/silcd/config"
	examplechain "github.com/silcprotocol/silcchain/silcd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	setupSDKConfig()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "silcd", examplechain.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}

func setupSDKConfig() {
	config := sdk.GetConfig()
	silcdconfig.SetBech32Prefixes(config)
	config.Seal()
}
