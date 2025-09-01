package formatx

import (
	"fmt"
)

// PrintTable prints a simple table for CLI output
func PrintTable(headers []string, rows [][]string) {
	for _, h := range headers {
		fmt.Printf("%-12s", h)
	}
	fmt.Println()
	for _, row := range rows {
		for _, col := range row {
			fmt.Printf("%-12s", col)
		}
		fmt.Println()
	}
}
