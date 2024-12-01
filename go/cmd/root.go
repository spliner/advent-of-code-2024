package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spliner/aoc2024/pkg/day1"
)

type solver func(*bufio.Scanner) (string, error)

func init() {
	addCmd(1, day1.Part1, day1.Part2)
}

var rootCmd = &cobra.Command{
	Use:     "aoc2024 [day] [input path] [part]",
	Short:   "Solutions for Advent of Code 2024",
	Version: "1.0.0",
	Example: "aoc2024 1 1 ../inputs/day1.txt",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addCmd(day int, part1, part2 solver) {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%d [part] [input path] ", day),
		Short: fmt.Sprintf("Day %d solution", day),
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			readFile, err := os.Open(args[1])
			if err != nil {
				return err
			}
			defer readFile.Close()

			scanner := bufio.NewScanner(readFile)
			scanner.Split(bufio.ScanLines)

			part := args[0]
			switch part {
			case "1":
				if part1 == nil {
					return fmt.Errorf("part 1 not yet implemented for day %d", day)
				}

				result, err := part1(scanner)
				if err != nil {
					return err
				}
				fmt.Printf("Day %d part 1: %s\n", day, result)
			case "2":
				if part2 == nil {
					return fmt.Errorf("part 2 not yet implemented for day %d", day)
				}

				result, err := part2(scanner)
				if err != nil {
					return err
				}
				fmt.Printf("Day %d part 2: %s\n", day, result)
			default:
				return fmt.Errorf("invalid part: %s", part)
			}

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}
