package day6_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/spliner/aoc2024/pkg/day5"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := ``

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day5.Part1(scanner)

	require.Nil(t, err)
	require.Equal(t, "", result)
}

func TestPart2(t *testing.T) {
	input := ``

	scanner := bufio.NewScanner(strings.NewReader(input))
	result, err := day5.Part2(scanner)

	require.Nil(t, err)
	require.Equal(t, "", result)
}
