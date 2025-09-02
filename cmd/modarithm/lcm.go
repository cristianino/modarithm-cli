package modarithm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

var (
	aLcm       int
	bLcm       int
	numbersLcm string
)

var lcmCmd = &cobra.Command{
	Use:   "lcm",
	Short: "Compute the least common multiple (LCM) of two or more integers.",
	Long: `Compute the least common multiple (LCM) of integers.
You can use either:
- Two specific integers: --a 12 --b 18
- Multiple integers: --numbers "12,18,24"`,
	Example: `modarithm lcm --a 12 --b 18
modarithm lcm --numbers "12,18,24"`,
	Run: func(cmd *cobra.Command, args []string) {
		if numbersLcm != "" {
			// Parse multiple numbers from the string
			numStrings := strings.Split(numbersLcm, ",")
			var numbers []int
			for _, numStr := range numStrings {
				numStr = strings.TrimSpace(numStr)
				num, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Printf("Error: invalid number '%s'\n", numStr)
					return
				}
				numbers = append(numbers, num)
			}

			if len(numbers) < 2 {
				fmt.Println("Error: need at least 2 numbers")
				return
			}

			result := mathx.LCMMultiple(numbers...)
			fmt.Printf("LCM(%s) = %d\n", strings.Join(numStrings, ", "), result)
		} else {
			// Use the original two-number approach
			result := mathx.LCM(aLcm, bLcm)
			fmt.Printf("LCM(%d, %d) = %d\n", aLcm, bLcm, result)
		}
	},
}

func init() {
	lcmCmd.Flags().IntVar(&aLcm, "a", 0, "First integer (required when not using --numbers)")
	lcmCmd.Flags().IntVar(&bLcm, "b", 0, "Second integer (required when not using --numbers)")
	lcmCmd.Flags().StringVar(&numbersLcm, "numbers", "", "Comma-separated list of numbers (e.g., \"12,18,24\")")

	// Make the flags mutually exclusive
	lcmCmd.MarkFlagsRequiredTogether("a", "b")
	lcmCmd.MarkFlagsMutuallyExclusive("numbers", "a")
	lcmCmd.MarkFlagsMutuallyExclusive("numbers", "b")

	rootCmd.AddCommand(lcmCmd)
}
