package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "keygent",
	Short: "Run commands with injected environment",
	Long: `Keygent is a lightweight environment runner.

It loads .env files or project configs and injects them into
your command execution safely and predictably.`,
	Example: `  keygent set TEST=test -- python app.py`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
