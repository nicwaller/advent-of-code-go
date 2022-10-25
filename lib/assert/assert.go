package assert

import "fmt"

func Equal[T comparable](actual T, expected T, name string) {
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
