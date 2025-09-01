package modarithm

import (
	"fmt"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

var (
	aLcm int
	bLcm int
)

var lcmCmd = &cobra.Command{
	Use:     "lcm",
	Short:   "Compute the least common multiple (LCM) of two integers.",
	Example: "modarithm lcm --a 12 --b 18",
	Run: func(cmd *cobra.Command, args []string) {
		result := mathx.LCM(aLcm, bLcm)
		fmt.Printf("LCM(%d, %d) = %d\n", aLcm, bLcm, result)
	},
}

func init() {
	lcmCmd.Flags().IntVar(&aLcm, "a", 0, "First integer (required)")
	lcmCmd.Flags().IntVar(&bLcm, "b", 0, "Second integer (required)")
	lcmCmd.MarkFlagRequired("a")
	lcmCmd.MarkFlagRequired("b")
	rootCmd.AddCommand(lcmCmd)
}
