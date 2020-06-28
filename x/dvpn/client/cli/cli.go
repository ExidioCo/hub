package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	deposit "github.com/sentinel-official/hub/x/dvpn/deposit/client/cli"
	node "github.com/sentinel-official/hub/x/dvpn/node/client/cli"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/client/cli"
	session "github.com/sentinel-official/hub/x/dvpn/session/client/cli"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/client/cli"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dvpn",
		Short: "Querying commands for the dVPN module",
	}

	cmd.AddCommand(deposit.GetQueryCommands(cdc)...)
	cmd.AddCommand(provider.GetQueryCommands(cdc)...)
	cmd.AddCommand(node.GetQueryCommands(cdc)...)
	cmd.AddCommand(subscription.GetQueryCommands(cdc)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dvpn",
		Short: "dVPN transactions subcommands",
	}

	cmd.AddCommand(provider.GetTxCommands(cdc)...)
	cmd.AddCommand(node.GetTxCommands(cdc)...)
	cmd.AddCommand(subscription.GetTxCommands(cdc)...)
	cmd.AddCommand(session.GetTxCommands(cdc)...)

	return cmd
}
