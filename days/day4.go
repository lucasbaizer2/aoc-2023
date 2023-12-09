package days

import (
	"fmt"
	"strings"

	"lucasbaizer.dev/aoc/2023/util"
)

func init() {
	addChallengeExecutor(4, 1, day4part1)
	addChallengeExecutor(4, 2, day4part2)
}

func day4part1(input string) string {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		split := strings.SplitN(strings.TrimSpace(line[strings.IndexRune(line, ':')+1:]), "|", 2)
		winning := util.SplitInts(split[0], " ")
		actual := util.SplitInts(split[1], " ")
		intersection := len(util.Intersection[int](winning, actual))
		if intersection > 0 {
			sum += 1 << (intersection - 1)
		}
	}
	return fmt.Sprintf("%d", sum)
}

type scratchcard struct {
	Number  int
	Winning []int
	Actual  []int
	Score   int
}

func day4part2(input string) string {
	lines := strings.Split(input, "\n")

	scratchcards := make([]scratchcard, 0, len(lines))
	for num, line := range lines {
		split := strings.SplitN(strings.TrimSpace(line[strings.IndexRune(line, ':')+1:]), "|", 2)
		winning := util.SplitInts(split[0], " ")
		actual := util.SplitInts(split[1], " ")
		intersection := len(util.Intersection[int](winning, actual))
		scratchcards = append(scratchcards, scratchcard{
			Number:  num,
			Winning: winning,
			Actual:  actual,
			Score:   intersection,
		})
	}

	cardQueue := util.NewDeque[scratchcard](len(scratchcards))
	for _, card := range scratchcards {
		cardQueue.Push(card)
	}

	sum := 0
	for cardQueue.Len() > 0 {
		sum += 1
		nextCard := cardQueue.Pop()
		if nextCard.Score > 0 {
			for num := 0; num < nextCard.Score; num++ {
				cardQueue.Push(scratchcards[nextCard.Number+num+1])
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}
