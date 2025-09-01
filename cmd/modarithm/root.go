package modarithm

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "modarithm",
	Short: "Modular arithmetic CLI toolkit for cryptography",
	Long:  "modarithm is a CLI tool for modular arithmetic, designed for cryptography and education.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	// Add all subcommands
	rootCmd.AddCommand(NewMultiplesCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
