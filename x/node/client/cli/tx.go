package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func txRegister(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			provider, err := hub.ProvAddressFromBech32(s)
			if err != nil {
				return err
			}

			s, err = cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			price, err := sdk.ParseCoins(s)
			if err != nil {
				return err
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegister(ctx.FromAddress, provider, price, remoteURL)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProvider, "", "node provider address")
	cmd.Flags().String(flagPrice, "", "node price per Gigabyte")
	cmd.Flags().String(flagRemoteURL, "", "node remove URL")

	_ = cmd.MarkFlagRequired(flagRemoteURL)

	return cmd
}

func txUpdate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			var provider hub.ProvAddress
			if len(s) > 0 {
				provider, err = hub.ProvAddressFromBech32(s)
				if err != nil {
					return err
				}
			}

			s, err = cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			var price sdk.Coins
			if len(s) > 0 {
				price, err = sdk.ParseCoins(s)
				if err != nil {
					return err
				}
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdate(ctx.FromAddress.Bytes(), provider, price, remoteURL)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProvider, "", "node provider address")
	cmd.Flags().String(flagPrice, "", "node price per Gigabyte")
	cmd.Flags().String(flagRemoteURL, "", "node remove URL")

	return cmd
}

func txSetStatus(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-set [Active | Inactive]",
		Short: "Set a node status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			msg := types.NewMsgSetStatus(ctx.FromAddress.Bytes(), hub.StatusFromString(args[0]))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}
