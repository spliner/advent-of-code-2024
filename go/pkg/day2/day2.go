package day2

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	var count int
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return "", err
		}

		line := scanner.Text()
		rawValues := strings.Split(line, " ")
		valid, err := isValid(rawValues)
		if err != nil {
			return "", err
		}

		if valid {
			count += 1
		}
	}

	return strconv.Itoa(count), nil
}

func isValid(rawValues []string) (bool, error) {
	var previousValue *int
	var previousDiff float64
	for _, rawValue := range rawValues {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return false, err
		}

		if previousValue != nil {
			diff := float64(value - *previousValue)
			absDiff := math.Abs(diff)
			if absDiff <= 0 || absDiff > 3 {
				return false, nil
			}

			if previousDiff < 0 && diff > 0 {
				return false, nil
			}

			if previousDiff > 0 && diff < 0 {
				return false, nil
			}

			previousDiff = diff
		}

		previousValue = &value
	}

	return true, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	var count int
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return "", err
		}

		line := scanner.Text()
		rawValues := strings.Split(line, " ")
		valid, err := isValid(rawValues)
		if err != nil {
			return "", err
		}

		if valid {
			count += 1
			continue
		}

		for i := range len(rawValues) {
			newValues := make([]string, len(rawValues))
			copy(newValues, rawValues)
			newValues = append(newValues[:i], newValues[i+1:]...)
			valid, err := isValid(newValues)
			if err != nil {
				return "", err
			}

			if valid {
				count += 1
				break
			}
		}
	}

	return strconv.Itoa(count), nil
}
