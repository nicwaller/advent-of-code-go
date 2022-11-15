package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Day(16)
	aoc.Test(run, "sample1.txt", "16", "")
	aoc.Test(run, "sample2.txt", "12", "")
	aoc.Test(run, "sample3.txt", "23", "")
	aoc.Test(run, "sample4.txt", "31", "")
	aoc.Test(run, "input.txt", "991", "1264485568252")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	var sb strings.Builder
	for _, c := range aoc.InputString() {
		sb.WriteString(fmt.Sprintf("%.4b", util.FromHexChar(byte(c))))
	}
	input := sb.String()
	p, _ := decodePacket(iter.StringIterator(input, 1))

	sumPacketVersions := 0
	depthFirstTraversal(p, func(p packet) {
		sumPacketVersions += p.version
	})
	*p1 = strconv.Itoa(sumPacketVersions)

	*p2 = strconv.Itoa(evalPacket(p))
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

func evalPacket(p packet) int {
	operands := f8l.Map[packet, int](p.sub, evalPacket)
	switch p.typeid {
	case 0: // sum
		return f8l.Reduce(operands, 0, util.IntSum)
	case 1: // product
		return f8l.Reduce(operands, 1, util.IntProduct)
	case 2: // min
		return f8l.Reduce(operands, math.MaxInt32, util.IntMin)
	case 3: // max
		return f8l.Reduce(operands, math.MinInt32, util.IntMax)
	case 4: // literal
		return p.operand
	case 5: // greater than
		if operands[0] > operands[1] {
			return 1
		} else {
			return 0
		}
	case 6: // less than
		if operands[0] < operands[1] {
			return 1
		} else {
			return 0
		}
	case 7: // equal to
		if operands[0] == operands[1] {
			return 1
		} else {
			return 0
		}
	default:
		panic(p.typeid)
	}
}
