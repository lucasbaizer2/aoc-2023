package days

import (
	"fmt"
	"strings"

	"lucasbaizer.dev/aoc/2023/util"
)

func init() {
	addChallengeExecutor(3, 1, day3part1)
	addChallengeExecutor(3, 2, day3part2)
}

type pos struct {
	x int
	y int
}

type part struct {
	num            int
	adjacentsStars []pos
}

func (p *part) getLength() int {
	return len(fmt.Sprintf("%d", p.num))
}

func day3part1(input string) string {
	lines := strings.Split(input, "\n")

	symbols := util.NewHashSet[pos](len(lines) * len(lines[0]))
	parts := make(map[pos]part, len(lines)*len(lines[0]))

	currentNumber := ""
	for y, line := range lines {
		line = strings.TrimSpace(line)

		// edge case where the number is the last thing on the line
		if len(currentNumber) > 0 {
			parts[pos{
				x: len(line) - len(currentNumber),
				y: y - 1,
			}] = part{
				num: util.MustParseInt(currentNumber),
			}
			currentNumber = ""
		}

		for x, ch := range line {
			if ch >= '0' && ch <= '9' {
				currentNumber += string(ch)
				continue
			} else {
				if len(currentNumber) > 0 {
					parts[pos{
						x - len(currentNumber),
						y,
					}] = part{
						num: util.MustParseInt(currentNumber),
					}
					currentNumber = ""
				}
			}
			if ch == '.' {
				continue
			}
			symbols.Add(pos{
				x,
				y,
			})
		}
	}

	sum := 0
partLoop:
	for partPos, part := range parts {
		plen := part.getLength()

		// top-left diagonal corner
		for x := max(0, partPos.x-1); x < partPos.x+plen+1; x++ {
			for y := max(0, partPos.y-1); y <= partPos.y+1; y++ {
				cpos := pos{
					x,
					y,
				}
				if symbols.Contains(cpos) {
					sum += part.num
					continue partLoop
				}
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func day3part2(input string) string {
	lines := strings.Split(input, "\n")

	symbols := make(map[pos]rune, len(lines)*len(lines[0]))
	parts := make(map[pos]*part, len(lines)*len(lines[0]))

	currentNumber := ""
	for y, line := range lines {
		line = strings.TrimSpace(line)

		// edge case where the number is the last thing on the line
		if len(currentNumber) > 0 {
			parts[pos{
				x: len(line) - len(currentNumber),
				y: y - 1,
			}] = &part{
				num:            util.MustParseInt(currentNumber),
				adjacentsStars: make([]pos, 0, (len(currentNumber)+2)*3),
			}
			currentNumber = ""
		}

		for x, ch := range line {
			if ch >= '0' && ch <= '9' {
				currentNumber += string(ch)
				continue
			} else {
				if len(currentNumber) > 0 {
					parts[pos{
						x - len(currentNumber),
						y,
					}] = &part{
						num:            util.MustParseInt(currentNumber),
						adjacentsStars: make([]pos, 0, (len(currentNumber)+2)*3),
					}
					currentNumber = ""
				}
			}
			if ch == '.' {
				continue
			}
			symbols[pos{
				x,
				y,
			}] = ch
		}
	}

	for partPos, part := range parts {
		plen := part.getLength()

		// top-left diagonal corner
		for x := max(0, partPos.x-1); x < partPos.x+plen+1; x++ {
			for y := max(0, partPos.y-1); y <= partPos.y+1; y++ {
				cpos := pos{
					x,
					y,
				}
				symbol, ok := symbols[cpos]
				if !ok {
					continue
				}
				if symbol == '*' {
					part.adjacentsStars = append(part.adjacentsStars, cpos)
				}
			}
		}
	}

	sum := 0
	for posA, partA := range parts {
		for posB, partB := range parts {
			if posA == posB {
				continue
			}
			for _, starA := range partA.adjacentsStars {
				for _, starB := range partB.adjacentsStars {
					if starA == starB {
						sum += partA.num * partB.num
					}
				}
			}
		}
	}

	// divide by 2 to account for each gear pair (A, B) being counted twice [as (B, A)]
	return fmt.Sprintf("%d", sum/2)
}
