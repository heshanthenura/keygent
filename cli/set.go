package cli

import (
	"fmt"

	"github.com/heshanthenura/keygent/internal/runner"
	"github.com/spf13/cobra"
)

var setCommand = &cobra.Command{
	Use:   "set",
	Short: "set - set env vars for a command",
	Long:  "set - set env vars for a command",
	RunE: func(cmd *cobra.Command, args []string) error {
		dashIndex := cmd.ArgsLenAtDash()

		if dashIndex == -1 {
			return fmt.Errorf("no separator provided (-- is missing)")
		}

		envVars := args[:dashIndex]
		commands := args[dashIndex:]

		return runner.Run(envVars, commands)
	},
}

func init() {
	rootCmd.AddCommand(setCommand)
}
