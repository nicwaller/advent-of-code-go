package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"fmt"
	"strings"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

// fromHexChar converts a hex character into its value and a success flag.
func fromHexChar(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	panic(c)
}

func parseFile() string {
	var sb strings.Builder
	for _, c := range aoc.InputString() {
		sb.WriteString(fmt.Sprintf("%.4b", fromHexChar(byte(c))))
	}
	return sb.String()
}

type packet struct {
	version int
	typeid  int
	operand int
	sub     []packet
}

const TypeLiteral = 4

func decodePacket(ite iter.Iterator[string]) (packet, error) {
	p := packet{}
	readBits := func(n int) (int, error) {
		sum := 0
		bits, err := ite.TakeArray(n)
		if err != nil {
			return -1, err
		}
		for i, bitStr := range bits {
			bit := 0
			if bitStr == "1" {
				bit = 1
			}
			sum |= bit << ((n - 1) - i)
		}
		return sum, nil
	}
	var err error
	p.version, err = readBits(3)
	if err != nil {
		return p, err
	}
	p.typeid, err = readBits(3)
	if err != nil {
		return p, err
	}
	if p.typeid == TypeLiteral {
		p.operand = 0
		for {
			keepGoing, _ := readBits(1)
			group, _ := readBits(4)
			p.operand = (p.operand << 4) | group
			if keepGoing == 0 {
				break
			}
		}
	} else {
		lengthType, _ := readBits(1)
		switch lengthType {
		case 0:
			maxBits, _ := readBits(15)
			payloadArr, _ := ite.TakeArray(maxBits)
			payloadStr := strings.Join(payloadArr, "")
			p.sub = decodeAllPackets(iter.StringIterator(payloadStr, 1))
		case 1:
			maxPackets, _ := readBits(11)
			for i := 0; i < maxPackets; i++ {
				innerPacket, _ := decodePacket(ite)
				//fmt.Println(innerPacket)
				p.sub = append(p.sub, innerPacket)
			}
		default:
			panic(fmt.Sprintf("invalid length type %v", lengthType))
		}
	}
	return p, nil
}

func decodeAllPackets(ite iter.Iterator[string]) []packet {
	ret := make([]packet, 0)
	for {
		last, err := decodePacket(ite)
		if err != nil {
			break
		}
		ret = append(ret, last)
	}
	return ret
}

func part1(input string) int {
	p, _ := decodePacket(iter.StringIterator(input, 1))
	sumPacketVersions := 0
	prettyPrintPacket(p, 0)
	depthFirstTraversal(p, func(p packet) {
		sumPacketVersions += p.version
	})
	return sumPacketVersions
}

func part2(input string) int {
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	return -1
}

func prettyPrintPacket(p packet, indent int) {
	fmt.Printf("%sV%d T%d (%d)\n", strings.Repeat("  ", indent), p.version, p.typeid, p.operand)
	for _, inner := range p.sub {
		prettyPrintPacket(inner, indent+1)
	}
}

// this is hard to implement as an iterator
func depthFirstTraversal(p packet, fn func(p packet)) {
	fn(p)
	for _, s := range p.sub {
		depthFirstTraversal(s, fn)
	}
}
