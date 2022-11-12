package f8l

func Map[I interface{}, O interface{}](items *[]I, mapFn func(I) O) []O {
	var results = make([]O, len(*items))
	for i, item := range *items {
		results[i] = mapFn(item)
	}
	return results
}

//func Atoi(items *[]string) []int {
//	return Map[string, int](items, util.UnsafeAtoi)
//}
//
//func Itoa(items *[]int) []string {
//	return Map[int, string](items, strconv.Itoa)
//}

func Reduce[T comparable](values *[]T, start T, reduce func(a T, b T) T) T {
	result := start
	for _, val := range *values {
		result = reduce(result, val)
	}
	return result
}

func Sum(values *[]int) int {
	add := func(a int, b int) int {
		return a + b
	}
	return Reduce[int](values, 0, add)
}

func Filter[T comparable](items *[]T, include func(T) bool) []T {
	var results = make([]T, 0)
	for _, item := range *items {
		if include(item) {
			results = append(results, item)
		}
	}
	return results
}
