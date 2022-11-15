package analyze

import (
	"golang.org/x/exp/constraints"
	"math"
)

func CountDistinct[T comparable](list []T) map[T]int {
	keys := make(map[T]int)
	for _, item := range list {
		//fmt.Println(item)
		if _, ok := keys[item]; !ok {
			keys[item] = 0
		}
		keys[item] += 1
	}
	return keys
}

func MostCommon[T comparable](list []T) T {
	distinct := CountDistinct(list)
	topScore := -1
	var topKey T
	for key := range distinct {
		occurrences := distinct[key]
		if occurrences > topScore {
			topScore = occurrences
			topKey = key
		}
	}
	return topKey
}

type AnalyzeResult[T comparable] struct {
	Min              T
	Max              T
	Count            int
	Distinct         int
	MostCommon       T
	CountMostCommon  int
	LeastCommon      T
	CountLeastCommon int
	Frequency        map[T]int
}

func Analyze[T constraints.Ordered](list []T) AnalyzeResult[T] {
	min := list[0]
	max := list[0]

	distinct := make(map[T]int)
	for _, item := range list {
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}

		if _, ok := distinct[item]; !ok {
			distinct[item] = 0
		}
		distinct[item] += 1
	}

	topScore := -1
	bottomScore := math.MaxInt32
	var topKey T
	var bottomKey T
	for key := range distinct {
		occurrences := distinct[key]
		if occurrences > topScore {
			topScore = occurrences
			topKey = key
		}
		if occurrences < bottomScore {
			bottomScore = occurrences
			bottomKey = key
		}
	}

	return AnalyzeResult[T]{
		Count:            len(list),
		Distinct:         len(distinct),
		Frequency:        distinct,
		Min:              min,
		Max:              max,
		MostCommon:       topKey,
		CountMostCommon:  topScore,
		LeastCommon:      bottomKey,
		CountLeastCommon: bottomScore,
	}
}

func MinMax[T constraints.Ordered](list []T) (T, T) {
	min := list[0]
	max := list[0]

	for _, item := range list {
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}

	return min, max
}
