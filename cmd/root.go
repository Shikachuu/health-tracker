package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "Start the health counter process",
	}

	cmd.AddCommand(NewCharacterCommand())
	cmd.AddCommand(NewGetMaxCommand())
	cmd.AddCommand(NewGetCurrentCommand())
	cmd.AddCommand(NewDamageCommand())
	cmd.AddCommand(NewHealCommand())

	return cmd
}
