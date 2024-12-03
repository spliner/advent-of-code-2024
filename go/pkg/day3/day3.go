package day3

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var regex *regexp.Regexp

func init() {
	regex = regexp.MustCompile(`mul\((\d+\,\d+)\)`)
}

func Part1(scanner *bufio.Scanner) (string, error) {
	input, err := mergeLines(scanner)
	if err != nil {
		return "", err
	}

	sum, err := process(input)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(sum), nil
}

func mergeLines(scanner *bufio.Scanner) (string, error) {
	var input string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}

		input += strings.TrimSpace(scanner.Text())
	}
	return input, nil
}

func process(input string) (int, error) {
	var sum int
	var count int
	matches := regex.FindAllStringSubmatch(input, -1)
	for _, m := range matches {
		split := strings.Split(m[1], ",")
		left, err := strconv.Atoi(string(split[0]))
		if err != nil {
			return 0, err
		}

		right, err := strconv.Atoi(string(split[1]))
		if err != nil {
			return 0, err
		}

		sum += left * right
		count += 1
	}

	return sum, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	input, err := mergeLines(scanner)
	if err != nil {
		return "", err
	}

	split := strings.Split(input, "don't")
	sum, err := process(split[0])
	if err != nil {
		return "", err
	}

	for _, dont := range split[1:] {
		_, after, found := strings.Cut(dont, "do()")
		if !found {
			continue
		}

		s, err := process(after)
		if err != nil {
			return "", err
		}

		sum += s
	}

	return strconv.Itoa(sum), nil
}
