package cli

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

// i have no idea why this is here, but i am afraid to touch it
var _ = strconv.Itoa(0)

func CmdAddTokensLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-tokens-lock \"[{\\\"amount\\\":\\\"2000\\\", \\\"denom\\\":\\\"doge\\\"}, {\\\"amount\\\":\\\"2\\\", \\\"denom\\\":\\\"btc\\\"},... ]\"",
		Short: "Broadcast message AddTokensLock",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAmountDenom := args[0]
			var coinsArg types.CoinsRaw
			err := json.Unmarshal([]byte(argsAmountDenom), &coinsArg)
			if err != nil {
				return errors.New("error while unmarshaling coins: " + err.Error())
			}

			coins, err := coinsArg.ParseCoins()
			if err != nil {
				return errors.New("error while unmarshaling coins: " + err.Error())
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddTokensLock(clientCtx.GetFromAddress().String(), coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
