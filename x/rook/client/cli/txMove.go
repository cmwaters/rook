package cli

// import (
  
// 	"github.com/spf13/cobra"

//     "github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	"github.com/cosmos/cosmos-sdk/client/tx"
// 	"github.com/cmwaters/rook/x/rook/types"
// )

// func CmdCreateMove() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "create-move [quantity] [] [to]",
// 		Short: "Creates a new move",
// 		Args:  cobra.ExactArgs(3),
// 		RunE: func(cmd *cobra.Command, args []string) error {
//       argsQuantity := string(args[0])
//       argsFrom := string(args[1])
//       argsTo := string(args[2])
      
// 			clientCtx := client.GetClientContextFromCmd(cmd)
// 			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgMove(clientCtx.GetFromAddress(), uint32(argsQuantity), string(argsFrom), string(argsTo))
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

//     return cmd
// }
