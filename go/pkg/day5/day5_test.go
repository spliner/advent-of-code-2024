package day5_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/spliner/aoc2024/pkg/day5"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day5.Part1(scanner)

	require.Nil(t, err)
	require.Equal(t, "143", result)
}

func TestPart2(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day5.Part2(scanner)

	require.Nil(t, err)
	require.Equal(t, "123", result)
}
