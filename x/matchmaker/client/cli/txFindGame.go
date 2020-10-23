package cli

import (
	"github.com/spf13/cobra"

	"github.com/cmwaters/rook/x/matchmaker/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateFindGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-findGame [game] [options]",
		Short: "Creates a new findGame",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGame := string(args[0])
			argsOptions := string(args[1])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgFindGame(clientCtx.GetFromAddress(), string(argsGame), string(argsOptions))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
