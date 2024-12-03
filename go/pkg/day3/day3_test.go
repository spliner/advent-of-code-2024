package day3_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/spliner/aoc2024/pkg/day3"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day3.Part1(scanner)

	require.Nil(t, err)
	require.Equal(t, "161", result)
}

func TestPart2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day3.Part2(scanner)

	require.Nil(t, err)
	require.Equal(t, "48", result)
}
