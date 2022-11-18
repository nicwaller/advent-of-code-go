package aoc

import (
	"fmt"
	"testing"
)

func TestBannerize(t *testing.T) {
	banner := Bannerize("BCEGHKLRYZ", "#")
	banner.Print()
	roundtrip := Debannerize(banner)
	fmt.Println(roundtrip)
}
