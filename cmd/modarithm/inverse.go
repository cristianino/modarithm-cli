package modarithm

import (
	"fmt"
	"os"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

var (
	aInv   int
	modInv int
)

var inverseCmd = &cobra.Command{
	Use:     "inverse",
	Short:   "Compute the modular inverse of a mod m, if it exists.",
	Example: "modarithm inverse --a 7 --mod 26",
	Run: func(cmd *cobra.Command, args []string) {
		inv, ok := mathx.ModInverse(aInv, modInv)
		if !ok {
			fmt.Fprintf(os.Stderr, "No modular inverse exists for %d mod %d\n", aInv, modInv)
			os.Exit(2)
		}
		fmt.Printf("Inverse of %d mod %d = %d\n", aInv, modInv, inv)
	},
}

func init() {
	inverseCmd.Flags().IntVar(&aInv, "a", 0, "Integer to invert (required)")
	inverseCmd.Flags().IntVar(&modInv, "mod", 0, "Modulus (required)")
	inverseCmd.MarkFlagRequired("a")
	inverseCmd.MarkFlagRequired("mod")
	rootCmd.AddCommand(inverseCmd)
}
