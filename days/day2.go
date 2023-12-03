package days

import (
	"fmt"
	"strings"

	"lucasbaizer.dev/aoc/2023/util"
)

func init() {
	addChallengeExecutor(2, 1, day2part1)
	addChallengeExecutor(2, 2, day2part2)
}

func day2part1(input string) string {
	const treds, tgreens, tblues int = 12, 13, 14

	getColorMax := func(color string) int {
		switch color {
		case "red":
			return treds
		case "green":
			return tgreens
		case "blue":
			return tblues
		default:
			fmt.Println(color)
			panic("invalid color")
		}
	}

	lines := strings.Split(input, "\n")
	sum := 0
lineLoop:
	for _, line := range lines {
		line = strings.TrimSpace(line)
		gameStr := line[5:strings.IndexRune(line, ':')]
		gameID := util.MustParseInt(gameStr)

		game := line[strings.IndexRune(line, ':')+2:]
		subgames := strings.Split(game, ";")
		for _, subgame := range subgames {
			subgame = strings.TrimSpace(subgame)
			choices := strings.Split(subgame, ",")
			for _, choice := range choices {
				choice = strings.TrimSpace(choice)
				choiceSplit := strings.Split(choice, " ")
				value := util.MustParseInt(choiceSplit[0])
				color := choiceSplit[1]
				if value > getColorMax(color) {
					continue lineLoop
				}
			}
		}

		sum += gameID
	}

	return fmt.Sprintf("%d", sum)
}

func day2part2(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		minRed, minGreen, minBlue := 0, 0, 0

		game := line[strings.IndexRune(line, ':')+2:]
		subgames := strings.Split(game, ";")
		for _, subgame := range subgames {
			subgame = strings.TrimSpace(subgame)
			choices := strings.Split(subgame, ",")
			for _, choice := range choices {
				choice = strings.TrimSpace(choice)
				choiceSplit := strings.Split(choice, " ")
				value := util.MustParseInt(choiceSplit[0])
				color := choiceSplit[1]
				if color == "red" && value > minRed {
					minRed = value
				} else if color == "green" && value > minGreen {
					minGreen = value
				} else if color == "blue" && value > minBlue {
					minBlue = value
				}
			}
		}

		sum += minRed * minGreen * minBlue
	}

	return fmt.Sprintf("%d", sum)
}
