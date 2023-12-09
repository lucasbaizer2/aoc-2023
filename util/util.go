package util

import (
	"strconv"
	"strings"
)

func SplitInts(s string, sep string) []int {
	split := strings.Split(s, sep)
	nums := make([]int, 0, len(split))
	for _, item := range split {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		nums = append(nums, MustParseInt(item))
	}
	return nums
}

func MustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func Intersection[T comparable](a []T, b []T) []T {
	intersection := make([]T, 0, min(len(a), len(b)))
	set := NewHashSet[T](min(len(a), len(b)))

	for _, v := range a {
		set.Add(v)
	}

	for _, v := range b {
		if set.Contains(v) {
			intersection = append(intersection, v)
		}
	}

	return intersection
}
