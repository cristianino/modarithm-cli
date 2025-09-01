package modarithm

import (
	"fmt"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

var (
	base   int
	exp    int
	modExp int
)

var modexpCmd = &cobra.Command{
	Use:     "modexp",
	Short:   "Compute base^exp mod m using fast modular exponentiation.",
	Example: "modarithm modexp --base 5 --exp 117 --mod 19",
	Run: func(cmd *cobra.Command, args []string) {
		result := mathx.ModExp(base, exp, modExp)
		fmt.Printf("%d^%d mod %d = %d\n", base, exp, modExp, result)
	},
}

func init() {
	modexpCmd.Flags().IntVar(&base, "base", 0, "Base integer (required)")
	modexpCmd.Flags().IntVar(&exp, "exp", 0, "Exponent (required)")
	modexpCmd.Flags().IntVar(&modExp, "mod", 0, "Modulus (required)")
	modexpCmd.MarkFlagRequired("base")
	modexpCmd.MarkFlagRequired("exp")
	modexpCmd.MarkFlagRequired("mod")
	rootCmd.AddCommand(modexpCmd)
}
