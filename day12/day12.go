package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/set"
	"fmt"
	"github.com/yourbasic/graph"
	"strings"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType struct {
	nodeNames  []string
	nodeLookup map[string]int
	g          *graph.Mutable
}

func parseFile() fileType {
	nodeSet := set.New[string]()
	edges := make([][2]string, 0)

	// Read the file into memory
	for line := aoc.InputLinesIterator(); line.Next(); {
		parts := strings.Split(line.Value(), "-")
		from := parts[0]
		to := parts[1]
		nodeSet.Extend(from, to)
		edges = append(edges, [2]string{from, to})
		edges = append(edges, [2]string{to, from})
	}

	// assign ordered numeric values to the nodeNames
	// I am surprised the graph library doesn't do this.
	nodeList := nodeSet.Items()
	nodeMap := make(map[string]int)
	for i, v := range nodeList {
		nodeMap[v] = i
		//fmt.Printf("  Node %d: %v\n", i, v)
	}

	// build the graph object
	g := graph.New(len(nodeList))
	for _, edge := range edges {
		g.Add(nodeMap[edge[0]], nodeMap[edge[1]])
	}
	//fmt.Println(g)

	return fileType{
		nodeNames:  nodeList,
		nodeLookup: nodeMap,
		g:          g,
	}
}

func isSmallCave(name string) bool {
	if name == "start" || name == "end" {
		return false
	}
	return strings.ToLower(name) == name
}

func getPathsFrom(v int, f fileType, pathSoFar []int, fn func(path []int)) {
	isSeen := func(nodeId int) bool {
		for _, v := range pathSoFar {
			if v == nodeId {
				return true
			}
		}
		return false
	}
	f.g.Visit(v, func(w int, c int64) bool {
		newPath := append(pathSoFar, w)
		nodeName := f.nodeNames[w]
		if nodeName == "start" {
			return false
		}
		if nodeName == "end" {
			fn(newPath)
			return false
		}
		if isSeen(w) && isSmallCave(f.nodeNames[w]) {
			//fmt.Println("dead end")
			return false
		}
		getPathsFrom(w, f, newPath, fn)
		return false
	})
}

func part1(input fileType) int {
	start := input.nodeLookup["start"]
	count := 0
	getPathsFrom(start, input, []int{start}, func(path []int) {
		count++
		//fmt.Printf("Got path: %v\n  ", path)
		//for _, node := range path {
		//	fmt.Printf("%s,", input.nodeNames[node])
		//}
		//fmt.Println("")
	})
	return count
}

func part2(input fileType) int {
	return -1
}
