package modarithm

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cristianino/modarithm-cli/internal/mathx"
	"github.com/spf13/cobra"
)

func NewMultiplesCmd() *cobra.Command {
	var (
		nStr     string
		limit    int
		limitNeg int
		limitPos int
		group    string
		asJSON   bool
	)

	cmd := &cobra.Command{
		Use:     "multiples",
		Short:   "Generate multiples of n (negatives, zero, positives)",
		Example: `modarithm multiples --n 5 --limit 4\nmodarithm multiples --n 5 --limitNeg 2 --limitPos 4\nmodarithm multiples --n 4,6 --limit 3 --group pos`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate mutually exclusive flags
			if limit > 0 && (limitNeg > 0 || limitPos > 0) {
				return errors.New("use either --limit or (--limitNeg and --limitPos), not both")
			}

			// Parse multiple n values
			nValues, err := parseNValues(nStr)
			if err != nil {
				return fmt.Errorf("invalid --n values: %v", err)
			}

			// Normalize and validate group
			g := normalizeGroup(group)
			if g == "" {
				return fmt.Errorf("invalid --group '%s': use neg|zero|pos|all", group)
			}

			// Generate multiples for each n value
			var results []mathx.Multiples
			for _, n := range nValues {
				var m mathx.Multiples
				var err error

				if limit > 0 {
					m, err = mathx.MultiplesSym(n, limit)
				} else {
					m, err = mathx.MultiplesAsym(n, limitNeg, limitPos)
				}
				if err != nil {
					return fmt.Errorf("error generating multiples for n=%d: %v", n, err)
				}
				results = append(results, m)
			}

			// Output results
			if len(results) == 1 {
				// Single n value - use existing logic
				m := results[0]
				return outputSingleResult(m, g, asJSON)
			} else {
				// Multiple n values - matrix output
				return outputMatrixResult(results, g, asJSON, nValues)
			}
		},
	}

	cmd.Flags().StringVar(&nStr, "n", "", "base integer(s) - single value or comma-separated (e.g. '5' or '4,6') (required)")
	cmd.MarkFlagRequired("n")

	cmd.Flags().IntVar(&limit, "limit", 0, "symmetric limit (both sides)")
	cmd.Flags().IntVar(&limitNeg, "limitNeg", 0, "asymmetric negative limit")
	cmd.Flags().IntVar(&limitPos, "limitPos", 0, "asymmetric positive limit")

	cmd.Flags().StringVar(&group, "group", "all", "group to show: neg|zero|pos|all")
	cmd.Flags().BoolVar(&asJSON, "json", false, "output as JSON")

	return cmd
}

func normalizeGroup(s string) mathx.Group {
	switch strings.ToLower(s) {
	case "neg", "negative", "negatives":
		return mathx.GroupNeg
	case "zero", "0":
		return mathx.GroupZero
	case "pos", "positive", "positives":
		return mathx.GroupPos
	case "all":
		return mathx.GroupAll
	default:
		return ""
	}
}

func parseNValues(nStr string) ([]int, error) {
	if nStr == "" {
		return nil, errors.New("n values cannot be empty")
	}

	parts := strings.Split(nStr, ",")
	var values []int

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid integer '%s': %v", part, err)
		}
		values = append(values, n)
	}

	if len(values) == 0 {
		return nil, errors.New("no valid n values found")
	}

	return values, nil
}

func outputSingleResult(m mathx.Multiples, g mathx.Group, asJSON bool) error {
	if asJSON {
		if g == mathx.GroupAll {
			// For 'all' group, provide both structured and flat array
			result := map[string]interface{}{
				"negatives": m.Negatives,
				"zero":      m.Zero,
				"positives": m.Positives,
				"all":       m.RenderGroup(mathx.GroupAll),
			}
			return encodeJSON(os.Stdout, result)
		}
		// For specific groups, return flat array
		return encodeJSON(os.Stdout, m.RenderGroup(g))
	}

	// Human-readable output
	switch g {
	case mathx.GroupNeg:
		fmt.Println(m.RenderGroup(g))
	case mathx.GroupZero:
		fmt.Println(0)
	case mathx.GroupPos:
		fmt.Println(m.RenderGroup(g))
	case mathx.GroupAll:
		fmt.Println(m.RenderGroup(g))
	default:
		// Fallback: show structured output
		fmt.Printf("Negatives: %v\nZero: %d\nPositives: %v\n", m.Negatives, m.Zero, m.Positives)
	}
	return nil
}

func outputMatrixResult(results []mathx.Multiples, g mathx.Group, asJSON bool, nValues []int) error {
	if asJSON {
		// JSON matrix output
		var matrix [][]int
		for _, m := range results {
			matrix = append(matrix, m.RenderGroup(g))
		}

		result := map[string]interface{}{
			"n_values": nValues,
			"group":    string(g),
			"matrix":   matrix,
		}
		return encodeJSON(os.Stdout, result)
	}

	// Human-readable matrix output
	for _, m := range results {
		fmt.Println(m.RenderGroup(g))
	}
	return nil
}

func encodeJSON(w *os.File, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
