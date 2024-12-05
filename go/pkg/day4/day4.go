package day4

import (
	"bufio"
	"strconv"
)

type point struct {
	r rune
	y int
	x int
}

type test []point

type step struct {
	x int
	y int
}

func Part1(scanner *bufio.Scanner) (string, error) {
	lines, err := parse(scanner)
	if err != nil {
		return "", err
	}

	var count int
	for y, line := range lines {
		for x, r := range line {
			if r == 'X' {
				tests := buildPart1Tests(x, y, []rune("XMAS"), lines)
				c := countMatches(lines, tests)
				count += c
			}
		}
	}

	return strconv.Itoa(count), nil
}

func parse(scanner *bufio.Scanner) ([][]rune, error) {
	lines := make([][]rune, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := []rune(scanner.Text())
		lines = append(lines, line)
	}

	return lines, nil
}

func buildPart1Tests(x, y int, word []rune, lines [][]rune) []test {
	buildTest := func(s step) test {
		cx, cy := x, y
		points := make([]point, len(word)-1)
		for i := 1; i < len(word); i++ {
			cx += s.x
			cy += s.y
			points[i-1] = point{word[i], cy, cx}
		}
		return points
	}

	steps := []step{
		{1, 0},   // ➡
		{-1, 0},  // ⬅
		{0, 1},   // ⬇
		{0, -1},  // ⬆
		{1, 1},   // ↘
		{1, -1},  // ↗
		{-1, 1},  // ↙
		{-1, -1}, // ↖
	}

	tests := make([]test, 0, 8)

	for _, s := range steps {
		minX, maxX, minY, maxY := x, x, y, y
		if s.x > 0 {
			maxX += s.x * (len(word) - 1)
		} else if s.x < 0 {
			minX += s.x * (len(word) - 1)
		}

		if s.y > 0 {
			maxY += s.y * (len(word) - 1)
		} else if s.y < 0 {
			minY += s.y * (len(word) - 1)
		}

		isInBounds := minX >= 0 && maxX < len(lines[y]) && minY >= 0 && maxY < len(lines)
		if isInBounds {
			tests = append(tests, buildTest(s))
		}
	}

	return tests
}

func Part2(scanner *bufio.Scanner) (string, error) {
	lines, err := parse(scanner)
	if err != nil {
		return "", err
	}

	var count int
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			if lines[y][x] == 'A' {
				tests := buildPart2Tests(x, y)
				c := countMatches(lines, tests)
				count += c
			}
		}
	}

	return strconv.Itoa(count), nil
}

func buildPart2Tests(x, y int) []test {
	return []test{
		{{'M', y - 1, x - 1}, {'S', y - 1, x + 1}, {'M', y + 1, x - 1}, {'S', y + 1, x + 1}},
		{{'S', y - 1, x - 1}, {'S', y - 1, x + 1}, {'M', y + 1, x - 1}, {'M', y + 1, x + 1}},
		{{'M', y - 1, x - 1}, {'M', y - 1, x + 1}, {'S', y + 1, x - 1}, {'S', y + 1, x + 1}},
		{{'S', y - 1, x - 1}, {'M', y - 1, x + 1}, {'S', y + 1, x - 1}, {'M', y + 1, x + 1}},
	}
}

func countMatches(lines [][]rune, tests []test) int {
	var count int
	for _, points := range tests {
		for i, p := range points {
			if lines[p.y][p.x] != p.r {
				break
			}

			if i == len(points)-1 {
				count++
			}
		}
	}

	return count
}
