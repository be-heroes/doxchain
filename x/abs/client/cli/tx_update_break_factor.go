package cli

import (
	"strconv"

	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateBreakFactor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-break-factor [breakFactor]",
		Short: "Broadcast message UpdateBreakFactor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			decValue, err := sdk.NewDecFromStr(args[0])

			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBreakFactorRequest(
				clientCtx.GetFromAddress().String(),
				decValue,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
