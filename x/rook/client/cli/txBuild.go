package cli

// import (

// 	"github.com/spf13/cobra"

//     "github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	"github.com/cosmos/cosmos-sdk/client/tx"
// 	"github.com/cmwaters/rook/x/rook/types"
// )

// func CmdCreateBuild() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "create-build [settlement] [position]",
// 		Short: "Creates a new build",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
//       argsSettlement := string(args[0])
//       argsPosition := string(args[1])

// 			clientCtx := client.GetClientContextFromCmd(cmd)
// 			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgBuild(clientCtx.GetFromAddress(), string(argsSettlement), string(argsPosition))
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

//     return cmd
// }
