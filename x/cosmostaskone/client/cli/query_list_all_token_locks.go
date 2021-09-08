package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

var _ = strconv.Itoa(0)

func CmdListAllTokenLocks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all-token-locks",
		Short: "Query ListAllTokenLocks",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListAllTokenLocksRequest{}

			res, err := queryClient.ListAllTokenLocks(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
