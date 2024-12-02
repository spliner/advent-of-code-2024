package day2_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/spliner/aoc2024/pkg/day2"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day2.Part1(scanner)

	require.Nil(t, err)
	require.Equal(t, "2", result)
}

func TestPart2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day2.Part2(scanner)

	require.Nil(t, err)
	require.Equal(t, "4", result)
}
