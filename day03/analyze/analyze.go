package analyze

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
