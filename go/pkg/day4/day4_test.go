package day4_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/spliner/aoc2024/pkg/day4"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day4.Part1(scanner)

	require.Nil(t, err)
	require.Equal(t, "18", result)
}

func TestPart2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day4.Part2(scanner)

	require.Nil(t, err)
	require.Equal(t, "9", result)
}
