package cli

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

var _ = strconv.Itoa(0)

type coinRaw struct {
	Amount string
	Denom  string
}

type coinsRaw []coinRaw

func (cr coinsRaw) ParseCoins() ([]*cosmosTypes.Coin, error) {
	var amount int64
	var err error
	coins := make([]*cosmosTypes.Coin, 0, len(cr))

	for _, v := range cr {
		coin := cosmosTypes.Coin{}
		amount, err = strconv.ParseInt(v.Amount, 10, 64)
		if err != nil {
			return nil, err
		}
		if v.Denom == "" {
			return nil, errors.New("error: no denom specified")
		}

		coin.Denom = v.Denom
		coin.Amount = cosmosTypes.NewInt(amount)
		coins = append(coins, &coin)
	}

	return coins, nil
}

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
