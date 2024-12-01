package day1

import (
	"bufio"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	lefts := make([]int, 0)
	rights := make([]int, 0)
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return "", err
		}

		l := strings.TrimSpace(scanner.Text())
		if l == "" {
			continue
		}

		parts := strings.Split(l, "   ")
		left, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return "", err
		}

		right, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return "", err
		}

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	var sum int
	for i := 0; i < len(lefts); i++ {
		left := lefts[i]
		right := rights[i]
		sum += int(math.Abs(float64(left) - float64(right)))
	}
	return strconv.Itoa(sum), nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	lefts := make([]int, 0)
	counts := make(map[int]int, 0)
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return "", err
		}

		l := strings.TrimSpace(scanner.Text())
		if l == "" {
			continue
		}

		parts := strings.Split(l, "   ")
		left, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return "", err
		}

		right, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return "", err
		}

		lefts = append(lefts, left)
		counts[right] += 1
	}

	var sum int
	for i := 0; i < len(lefts); i++ {
		left := lefts[i]
		count := counts[left]
		sum += left * count
	}
	return strconv.Itoa(sum), nil
}
