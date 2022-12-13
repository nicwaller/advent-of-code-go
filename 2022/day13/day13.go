package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/stack"
	"advent-of-code/lib/util"
	"golang.org/x/exp/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 13)
	aoc.Test(run, "sample.txt", "13", "140")
	//aoc.Test(run, "input.txt", "6070", "20758")
	aoc.Run(run)
	aoc.Out()
}

type Thing struct {
	children []*Thing
	value    int
	isList   bool
	id       int
}

func NewListThing() Thing {
	return Thing{
		isList:   true,
		children: make([]*Thing, 0),
		id:       rand.Int(),
	}
}

func NewValueThing(v int) Thing {
	return Thing{
		isList: false,
		value:  v,
	}
}

func (t *Thing) AddValue(v int) {
	n := NewValueThing(v)
	t.children = append(t.children, &n)
}

func (t *Thing) String() string {
	if t.isList {
		return "[" + strings.Join(f8l.Map(t.children, func(i *Thing) string {
			return i.String()
		}), ",") + "]"
	} else {
		return strconv.Itoa(t.value)
	}
}

func tokenize(s string) []string {
	var intMatcher = regexp.MustCompile("[0-9]+")
	ret := make([]string, 0)
	i := 0
	for i < len(s) {
		switch s[i] {
		case '[':
			fallthrough
		case ']':
			fallthrough
		case ',':
			ret = append(ret, s[i:i+1])
			i++
		default:
			m := intMatcher.FindString(s[i:])
			ret = append(ret, m)
			i += len(m)
		}
	}
	return ret
}

func parse(s string) Thing {
	stack := stack.NewStack[*Thing]()
	root := NewListThing()
	stack.Push(&root)
	for _, tok := range tokenize(s) {
		switch tok {
		case "[":
			parent := stack.Peek()
			child := NewListThing()
			parent.children = append(parent.children, &child)
			stack.Push(&child)
		case "]":
			stack.MustPop()
		case ",":
			// okay?
		default:
			// a number!
			t := stack.Peek()
			t.AddValue(util.UnsafeAtoi(tok))
		}
	}
	return *root.children[0]
}

func cmp(t1 Thing, t2 Thing) int {
	if !t1.isList && !t2.isList {
		// both are integers
		if t1.value < t2.value {
			return -1
		} else if t1.value == t2.value {
			return 0
		} else {
			return 1
		}
	} else if t1.isList && t2.isList {
		// both values are lists
		mm := util.IntMin(len(t1.children), len(t2.children))
		for i := 0; i < mm; i++ {
			c := cmp(*t1.children[i], *t2.children[i])
			if c != 0 {
				return c
			}
		}
		if len(t1.children) < len(t2.children) {
			return -1
		} else if len(t1.children) > len(t2.children) {
			return 1
		} else {
			//panic("continue checking the next part of the input?!")
			return 0
		}
	} else if !t1.isList && t2.isList {
		sub := NewListThing()
		sub.AddValue(t1.value)
		return cmp(sub, t2)
	} else if t1.isList && !t2.isList {
		sub := NewListThing()
		sub.AddValue(t2.value)
		return cmp(t1, sub)
	} else {
		panic("what is this?!")
	}

}

func run(p1 *string, p2 *string) {
	inOrderSum := 0
	pEnumerator := iter.Enumerate(aoc.ParagraphsIterator())
	for pEnumerator.Next() {
		i := pEnumerator.Value().Index
		l1, l2 := util.Pair(pEnumerator.Value().Value)
		cmpResult := cmp(parse(l1), parse(l2))
		if cmpResult == -1 {
			inOrderSum += i + 1
		}
	}
	*p1 = strconv.Itoa(inOrderSum)

	packets := make([]Thing, 0)
	marker1 := parse("[[2]]")
	marker2 := parse("[[6]]")
	packets = append(packets, marker1, marker2)
	pIter := aoc.ParagraphsIterator()
	for pIter.Next() {
		l1, l2 := util.Pair(pIter.Value())
		packets = append(packets, parse(l1), parse(l2))
	}
	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) == -1
	})
	var y int
	var z int
	for i, v := range packets {
		if cmp(v, marker1) == 0 {
			y = i + 1
		}
		if cmp(v, marker2) == 0 {
			z = i + 1
		}
		//fmt.Printf("P[%d]: %v\n", i, v.String())
	}
	*p2 = strconv.Itoa(y * z)
}
