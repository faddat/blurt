package cli

import (
  
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/faddat/blurt/x/blurt/types"
)

func CmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-post [title] [body] [ipfs] [parent] [category]",
		Short: "Creates a new post",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsTitle := string(args[0])
      argsBody := string(args[1])
      argsIpfs := string(args[2])
      argsParent := string(args[3])
      argsCategory := string(args[4])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(clientCtx.GetFromAddress().String(), string(argsTitle), string(argsBody), string(argsIpfs), string(argsParent), string(argsCategory))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-post [id] [title] [body] [ipfs] [parent] [category]",
		Short: "Update a post",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]
      argsTitle := string(args[1])
      argsBody := string(args[2])
      argsIpfs := string(args[3])
      argsParent := string(args[4])
      argsCategory := string(args[5])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePost(clientCtx.GetFromAddress().String(), id, string(argsTitle), string(argsBody), string(argsIpfs), string(argsParent), string(argsCategory))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeletePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-post [id] [title] [body] [ipfs] [parent] [category]",
		Short: "Delete a post by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]

        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePost(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
