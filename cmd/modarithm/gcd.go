package modarithm

import (
	"fmt"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

var (
	a int
	b int
)

var gcdCmd = &cobra.Command{
	Use:     "gcd",
	Short:   "Compute the greatest common divisor (GCD) of two integers.",
	Example: "modarithm gcd --a 84 --b 30",
	Run: func(cmd *cobra.Command, args []string) {
		result := mathx.GCD(a, b)
		fmt.Printf("GCD(%d, %d) = %d\n", a, b, result)
	},
}

func init() {
	gcdCmd.Flags().IntVar(&a, "a", 0, "First integer (required)")
	gcdCmd.Flags().IntVar(&b, "b", 0, "Second integer (required)")
	gcdCmd.MarkFlagRequired("a")
	gcdCmd.MarkFlagRequired("b")
	rootCmd.AddCommand(gcdCmd)
}
