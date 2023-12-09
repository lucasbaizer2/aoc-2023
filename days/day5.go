package days

import (
	"fmt"
	"sort"
	"strings"

	"lucasbaizer.dev/aoc/2023/util"
)

func init() {
	addChallengeExecutor(5, 1, day5part1)
	addChallengeExecutor(5, 2, day5part2)
}

type almanacRange struct {
	Destination int
	Source      int
	Length      int
}

type almanacMap struct {
	Target string
	Ranges []*almanacRange
}

func (m *almanacMap) MapValue(val int) int {
	for _, r := range m.Ranges {
		if val >= r.Source && val <= r.Source+r.Length {
			return r.Destination + (val - r.Source)
		}
	}
	return val
}

func (m *almanacMap) MapRange(seeds seedRange) []seedRange {
	ranges := make([]seedRange, 0, len(m.Ranges))
	for _, r := range m.Ranges {
		if seeds.Start > r.Source+r.Length || r.Source > seeds.End {
			continue
		}
		overlapStart := max(r.Source, seeds.Start)
		overlapEnd := min(r.Source+r.Length, seeds.End)
		mapped := seedRange{
			Start: r.Destination + (overlapStart - r.Source),
			End:   r.Destination + (overlapStart - r.Source) + (overlapEnd - overlapStart),
		}
		fmt.Printf("  [%d, %d] & [%d, %d] == [%d, %d] => [%d, %d]\n",
			seeds.Start, seeds.End,
			r.Source, r.Source+r.Length,
			overlapStart, overlapEnd,
			mapped.Start, mapped.End,
		)
		ranges = append(ranges, mapped)
	}
	return ranges
}

type almanac struct {
	Seeds []int
	Maps  map[string]*almanacMap
}

func parseAlmanac(input string) *almanac {
	sections := strings.Split(input, "\n\n")
	seeds := util.SplitInts(sections[0][len("seeds: "):], " ")
	mapSections := sections[1:]

	maps := make(map[string]*almanacMap, len(mapSections))
	for _, mapSection := range mapSections {
		lines := strings.Split(mapSection, "\n")
		fromTo := strings.Split(lines[0][:strings.IndexRune(lines[0], ' ')], "-to-")
		from := fromTo[0]
		to := fromTo[1]

		ranges := make([]*almanacRange, 0, len(lines)-1)
		for _, line := range lines[1:] {
			vals := util.SplitInts(line, " ")
			ranges = append(ranges, &almanacRange{
				Destination: vals[0],
				Source:      vals[1],
				Length:      vals[2],
			})
		}

		maps[from] = &almanacMap{
			Target: to,
			Ranges: ranges,
		}
	}

	return &almanac{
		Seeds: seeds,
		Maps:  maps,
	}
}

func day5part1(input string) string {
	almanac := parseAlmanac(input)

	locations := make([]int, 0, len(almanac.Seeds))
	for _, seed := range almanac.Seeds {
		currentVal := seed
		currentMap := almanac.Maps["seed"]
		for {
			currentVal = currentMap.MapValue(currentVal)
			nextMap, ok := almanac.Maps[currentMap.Target]
			if !ok {
				break
			}
			currentMap = nextMap
		}
		locations = append(locations, currentVal)
	}

	sort.Ints(locations)

	return fmt.Sprintf("%d", locations[0])
}

type seedRange struct {
	Start int
	End   int
}

func day5part2(input string) string {
	almanac := parseAlmanac(input)

	seeds := make([]seedRange, 0, 1024)
	for i := 0; i < len(almanac.Seeds); i += 2 {
		start := almanac.Seeds[i]
		end := start + almanac.Seeds[i+1] - 1
		seeds = append(seeds, seedRange{
			Start: start,
			End:   end,
		})
	}

	locations := make([]seedRange, 0, len(seeds))
	for _, seed := range seeds {
		currentVals := []seedRange{seed}
		currentMap := almanac.Maps["seed"]
		for {
			mappedVals := make([]seedRange, 0, 10)
			for _, val := range currentVals {
				mappedVals = append(mappedVals, currentMap.MapRange(val)...)
			}
			currentVals = mappedVals
			nextMap, ok := almanac.Maps[currentMap.Target]
			if !ok {
				locations = append(locations, currentVals...)
				break
			}
			currentMap = nextMap
		}
	}

	sort.Slice(locations, func(i, j int) bool {
		a := &locations[i]
		b := &locations[j]
		return a.Start < b.Start
	})

	return fmt.Sprintf("%d", locations[0].Start)
}
