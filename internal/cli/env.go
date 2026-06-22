package cli

import (
	"fmt"

	parser "github.com/heshanthenura/keygent/internal/env"
	"github.com/heshanthenura/keygent/internal/runner"
	"github.com/spf13/cobra"
)

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "env - use env file to set env vars for a command",
	Long:  "env - use env file to set env vars for a command",
	RunE: func(cmd *cobra.Command, args []string) error {

		dashIndex := cmd.ArgsLenAtDash()
		if dashIndex == -1 {
			return fmt.Errorf("no separator provided (-- is missing)")
		}

		if dashIndex == 0 {
			return fmt.Errorf("missing env file path")
		}

		envFile := args[0]

		envVars, err := parser.ParseEnvFile(envFile)
		if err != nil {
			return err
		}

		commands := args[dashIndex:]
		if len(commands) == 0 {
			return fmt.Errorf("no command provided after --")
		}

		return runner.Run(envVars, commands)
	},
}

func init() {
	rootCmd.AddCommand(envCommand)
}
