package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 7)
	aoc.Test(run, "sample.txt", "95437", "24933642")
	aoc.Test(run, "input.txt", "1723892", "8474158")
	aoc.Run(run)
	aoc.Out()
}

type fsentry struct {
	name     string
	size     int
	parent   *fsentry
	children []*fsentry
}

func makefile(name string, size int) *fsentry {
	return &fsentry{
		name:     name,
		size:     size,
		children: nil,
	}
}

func makedir(name string) *fsentry {
	return &fsentry{
		name:     name,
		size:     0,
		children: make([]*fsentry, 0),
	}
}

func (fs *fsentry) addChild(e *fsentry) *fsentry {
	e.parent = fs
	fs.children = append(fs.children, e)
	return e
}

func (fs *fsentry) absolutePath() []string {
	if fs.parent == nil {
		return []string{"root"}
	} else {
		return append(fs.parent.absolutePath(), fs.name)
	}
}

func (fs *fsentry) absolutePathStr() string {
	return strings.Join(fs.absolutePath(), "/")
}

func (fs *fsentry) totalsize() int {
	// TODO: memoize this
	if fs.children != nil {
		sizes := f8l.Map[*fsentry, int](fs.children, func(e *fsentry) int {
			return e.totalsize()
		})
		return f8l.Sum(sizes)
	} else {
		return fs.size
	}
}

func run(p1 *string, p2 *string) {
	root := makedir("/")
	wd := root
	lines := aoc.InputLines()[1:]
	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			ff := strings.Fields(line)
			switch ff[1] {
			case "cd":
				if ff[2] == ".." {
					wd = wd.parent
				} else {
					subdir := makedir(ff[2])
					wd.addChild(subdir)
					wd = subdir
				}
			case "ls":
				// pass
			default:
				panic(ff[1])
			}
		} else {
			sizeStr, name := util.Pair(strings.Fields(line))
			if sizeStr == "dir" {
				// TODO: create dir maybe?
			} else {
				size := util.UnsafeAtoi(sizeStr)
				wd.addChild(makefile(name, size))
			}
		}
	}

	// Find all of the directories with a total size of at most 100000.
	// What is the sum of the total sizes of those directories?
	orderedSizes := make([]int, 0)
	sum := 0
	walkTree(root, func(e *fsentry) {
		if e.size == 0 {
			//fmt.Printf("%s %d\n", e.absolutePathStr(), e.totalsize())
			tot := e.totalsize()
			orderedSizes = append(orderedSizes, tot)
			if tot <= 100000 {
				sum += tot
			}
		}
	})
	*p1 = strconv.Itoa(sum)

	totalDisk := 70000000
	required := 30000000
	used := root.totalsize()
	mustBeRecovered := used - (totalDisk - required)
	sort.Ints(orderedSizes)
	for _, x := range orderedSizes {
		if x >= mustBeRecovered {
			*p2 = strconv.Itoa(x)
			break
		}
	}

}

func walkTree(root *fsentry, cb func(e *fsentry)) {
	cb(root)
	if root.children != nil {
		for _, c := range root.children {
			walkTree(c, cb)
		}
	}
}
