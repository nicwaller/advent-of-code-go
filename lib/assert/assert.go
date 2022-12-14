package assert

import "fmt"

func NotEqual[T comparable](actual T, expected T) {
	if actual != expected {
		return
	}
	panic(fmt.Sprintf("Failed Assertion: Actual %v == Expected %v",
		actual, expected))
}

func Equal[T comparable](actual T, expected T) {
	if actual == expected {
		return
	}
	panic(fmt.Sprintf("Failed Assertion: Actual %v != Expected %v",
		actual, expected))
}

func EqualNamed[T comparable](actual T, expected T, name string) {
	if actual == expected {
		return
	}
	panic(fmt.Sprintf("Failed Assertion: %s is %v (expected %v)",
		name, actual, expected))
}

func EqualAny[T comparable](actual T, expected []T, name string) {
	for _, exp := range expected {
		if actual == exp {
			return
		}
	}
	panic(fmt.Sprintf("Failed Assertion: %s is %v (expected one of %v)",
		name, actual, expected))
}
