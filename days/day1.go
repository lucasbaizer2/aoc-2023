package days

import (
	"fmt"
	"regexp"
	"strings"

	"lucasbaizer.dev/aoc/2023/util"
)

func init() {
	addChallengeExecutor(1, 1, day1part1)
	addChallengeExecutor(1, 2, day1part2)
}

func day1part1(input string) string {
	pattern := regexp.MustCompile(`^[a-z]*([0-9]).*([0-9])[a-z]*$`)
	backup := regexp.MustCompile(`^.*([0-9]).*$`)

	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if pattern.MatchString(line) {
			matches := pattern.FindStringSubmatch(line)
			if len(matches) == 3 {
				sum += util.MustParseInt(matches[1] + matches[2])
			} else {
				panic("matches != 3")
			}
		} else if backup.MatchString(line) {
			matches := backup.FindStringSubmatch(line)
			if len(matches) == 2 {
				sum += util.MustParseInt(matches[1] + matches[1])
			} else {
				panic("matches != 2")
			}
		} else {
			fmt.Println(line)
			panic("unmatchable line")
		}
	}

	return fmt.Sprintf("%d", sum)
}

func day1part2(input string) string {
	pattern := regexp.MustCompile(`^[a-z]*?([0-9]|zero|one|two|three|four|five|six|seven|eight|nine).*([0-9]|zero|one|two|three|four|five|six|seven|eight|nine)[a-z]*$`)
	backup := regexp.MustCompile(`^.*?([0-9]|zero|one|two|three|four|five|six|seven|eight|nine).*$`)

	lines := strings.Split(input, "\n")

	mapNum := func(s string) string {
		switch s {
		case "zero":
			return "0"
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		case "four":
			return "4"
		case "five":
			return "5"
		case "six":
			return "6"
		case "seven":
			return "7"
		case "eight":
			return "8"
		case "nine":
			return "9"
		}

		return s
	}

	sum := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if pattern.MatchString(line) {
			matches := pattern.FindStringSubmatch(line)
			if len(matches) == 3 {
				sum += util.MustParseInt(mapNum(matches[1]) + mapNum(matches[2]))
			} else {
				panic("matches != 3")
			}
		} else if backup.MatchString(line) {
			matches := backup.FindStringSubmatch(line)
			if len(matches) == 2 {
				sum += util.MustParseInt(mapNum(matches[1]) + mapNum(matches[1]))
			} else {
				panic("matches != 2")
			}
		} else {
			fmt.Println(line)
			panic("unmatchable line")
		}
	}

	return fmt.Sprintf("%d", sum)
}
