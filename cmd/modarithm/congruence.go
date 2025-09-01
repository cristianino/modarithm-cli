package modarithm

import (
	"fmt"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

var (
	aCon   int
	bCon   int
	modCon int
)

var congruenceCmd = &cobra.Command{
	Use:     "congruence",
	Short:   "Check if a ≡ b (mod m) and show canonical residues.",
	Example: "modarithm congruence --a 45 --b 9 --mod 6",
	Run: func(cmd *cobra.Command, args []string) {
		eq, ra, rb := mathx.Congruence(aCon, bCon, modCon)
		fmt.Printf("%d ≡ %d (mod %d)? %v\nCanonical residues: %d, %d\n", aCon, bCon, modCon, eq, ra, rb)
	},
}

func init() {
	congruenceCmd.Flags().IntVar(&aCon, "a", 0, "First integer (required)")
	congruenceCmd.Flags().IntVar(&bCon, "b", 0, "Second integer (required)")
	congruenceCmd.Flags().IntVar(&modCon, "mod", 0, "Modulus (required)")
	congruenceCmd.MarkFlagRequired("a")
	congruenceCmd.MarkFlagRequired("b")
	congruenceCmd.MarkFlagRequired("mod")
	rootCmd.AddCommand(congruenceCmd)
}
