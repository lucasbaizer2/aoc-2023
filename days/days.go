package days

import "fmt"

type ChallengeID struct {
	Day  int
	Part int
}

type ChallengeExecutor func(input string) string

var challenges map[ChallengeID]ChallengeExecutor = make(map[ChallengeID]ChallengeExecutor, 25*2)

func addChallengeExecutor(day, part int, challenge ChallengeExecutor) {
	challenges[ChallengeID{
		Day:  day,
		Part: part,
	}] = challenge
}

func GetChallengeExecutor(day, part int) (ChallengeExecutor, error) {
	challenge, ok := challenges[ChallengeID{
		Day:  day,
		Part: part,
	}]
	if !ok {
		return nil, fmt.Errorf("no implementation provided for day=%d part=%d", day, part)
	}
	return challenge, nil
}
