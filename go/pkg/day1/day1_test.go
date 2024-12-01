package day1_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/spliner/aoc2024/pkg/day1"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day1.Part1(scanner)

	require.Nil(t, err)
	require.Equal(t, result, "11")
}

func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day1.Part2(scanner)

	require.Nil(t, err)
	require.Equal(t, result, "31")
}
