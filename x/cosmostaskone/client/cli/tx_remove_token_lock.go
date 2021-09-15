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

func CmdRemoveTokenLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-token-lock [id]",
		Short: "Broadcast message RemoveTokenLock",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveTokenLock(clientCtx.GetFromAddress().String(), string(argsId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
