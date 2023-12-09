package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"lucasbaizer.dev/aoc/2023/days"
)

func main() {
	cmd := &cobra.Command{
		Use:   "aoc-2023",
		Short: "Advent of Code 2023",
		RunE: func(cmd *cobra.Command, args []string) error {
			day, err := cmd.PersistentFlags().GetInt("day")
			if err != nil {
				return err
			}
			part, err := cmd.PersistentFlags().GetInt("part")
			if err != nil {
				return err
			}
			input, err := cmd.PersistentFlags().GetString("input")
			if err != nil {
				return err
			}
			if input != "practice" && input != "real" {
				return fmt.Errorf("input must be either 'real' or 'practice', got '%s'", input)
			}

			inputTxt, err := os.ReadFile(fmt.Sprintf("inputs/day%d/%s.txt", day, input))
			if err != nil {
				fmt.Fprintf(os.Stderr, "input file '%s.txt' does not exist for day %d\n", input, day)
				os.Exit(1)
			}

			challenge, err := days.GetChallengeExecutor(day, part)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			before := time.Now()
			result := challenge(string(inputTxt))

			fmt.Printf("Challenge result: '%s' (took %dms)\n", result, time.Since(before).Milliseconds())

			return nil
		},
		DisableAutoGenTag: true,
	}

	cmd.PersistentFlags().IntP("day", "d", 1, "challenge day (1-25 inclusive)")
	cmd.PersistentFlags().IntP("part", "p", 1, "challenge part (either 1 or 2)")
	cmd.PersistentFlags().StringP("input", "i", "real", "challenge input type ('real' or 'practice')")
	cmd.MarkPersistentFlagRequired("day")
	cmd.MarkPersistentFlagRequired("part")
	cmd.MarkPersistentFlagRequired("input")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
