package binary

// maybe I should just use https://pkg.go.dev/math/bits

func NthBit(n int, value int) bool {
	// 0th bit = least siginficant bit
	// 01101001 <- value
	// 76543210 <- bit positions
	mask := 1 << n
	return (value & mask) == 1
}

func NthBitI(n int, value int) int {
	mask := 1 << n
	return (value & mask) >> n
}
