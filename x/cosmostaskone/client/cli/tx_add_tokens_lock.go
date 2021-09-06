package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

var _ = strconv.Itoa(0)

func CmdAddTokensLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-tokens-lock [amount] [denom] [address]",
		Short: "Broadcast message AddTokensLock",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAmount := string(args[0])
			argsDenom := string(args[1])
			argsAddress := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddTokensLock(clientCtx.GetFromAddress().String(), string(argsAmount), string(argsDenom), string(argsAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
