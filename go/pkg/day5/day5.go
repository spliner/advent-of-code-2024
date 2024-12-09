package day5

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	rules, err := parseRules(scanner)
	if err != nil {
		return "", err
	}

	var sum int
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}

		line := strings.TrimSpace(scanner.Text())
		pages := strings.Split(line, ",")
		valid := isValid(pages, rules)

		if valid {
			n, err := strconv.Atoi(pages[len(pages)/2])
			if err != nil {
				return "", err
			}

			sum += n
		}
	}

	return strconv.Itoa(sum), nil
}

func parseRules(scanner *bufio.Scanner) (map[string]map[string]struct{}, error) {
	rules := make(map[string]map[string]struct{})
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		before, after, _ := strings.Cut(line, "|")
		pages, ok := rules[before]
		if !ok {
			pages = make(map[string]struct{})
		}
		pages[after] = struct{}{}
		rules[before] = pages
	}

	return rules, nil
}

func isValid(pages []string, rules map[string]map[string]struct{}) bool {
	for i := 1; i < len(pages); i++ {
		page := pages[i]
		pageRules, ok := rules[page]
		if !ok {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			other := pages[j]
			_, ok := pageRules[other]
			if ok {
				return false
			}
		}
	}

	return true
}

func Part2(scanner *bufio.Scanner) (string, error) {
	rules, err := parseRules(scanner)
	if err != nil {
		return "", err
	}

	var sum int
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}

		line := strings.TrimSpace(scanner.Text())
		pages := strings.Split(line, ",")
		valid := isValid(pages, rules)

		if !valid {
			slices.SortFunc(pages, func(a, b string) int {
				pageRules, ok := rules[a]
				if !ok {
					return 0
				}

				_, ok = pageRules[b]
				if ok {
					return -1
				}

				return 1
			})
			n, err := strconv.Atoi(pages[len(pages)/2])
			if err != nil {
				return "", err
			}

			sum += n
		}
	}

	return strconv.Itoa(sum), nil
}
