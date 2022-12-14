package grapheasy

import (
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/stack"
	"fmt"
	"github.com/yourbasic/graph"
)

// add named nodes
// and edges between named nodes
// but still get the underlying

type Graph[T any] struct {
	underlying  *graph.Mutable
	nodeNames   []string
	nodeContext []T
	nodeIds     map[string]int
}

func New[T any](order int) Graph[T] {
	return Graph[T]{
		underlying:  graph.New(order),
		nodeNames:   make([]string, 0),
		nodeContext: make([]T, 0),
		nodeIds:     make(map[string]int),
	}
}

func (g *Graph[T]) Underlying() *graph.Mutable {
	return g.underlying
}

func (g *Graph[T]) DefineNodeContext(fn func(idx int, label string) T) {
	if len(g.nodeNames) != len(g.nodeContext) {
		panic("vector length mismatch")
	}
	for i := 0; i < len(g.nodeContext); i++ {
		label := g.nodeNames[i]
		g.nodeContext[i] = fn(i, label)
	}
}

func (g *Graph[T]) AddContext(label string, ctx T) {
	id := g.node(label)
	g.nodeNames[id] = label
	g.nodeContext[id] = ctx
}

func (g *Graph[T]) Add(from string, to string) {
	g.underlying.Add(g.node(from), g.node(to))
}

func (g *Graph[T]) AddBoth(label1 string, label2 string) {
	g.underlying.AddBoth(g.node(label1), g.node(label2))
}

func (g *Graph[T]) AddCost(from string, to string, cost int) {
	g.underlying.AddCost(g.node(from), g.node(to), int64(cost))
}

func (g *Graph[T]) AddBothCost(from string, to string, cost int) {
	g.underlying.AddCost(g.node(from), g.node(to), int64(cost))
	g.underlying.AddCost(g.node(to), g.node(from), int64(cost))
}

func (g *Graph[T]) NodeById(id int) (int, *string, *T) {
	return id, &g.nodeNames[id], &g.nodeContext[id]
}

func (g *Graph[T]) NodeByName(label string) (int, *string, *T) {
	return g.NodeById(g.node(label))
}

func (g *Graph[T]) Cost(node1 string, node2 string) int {
	n1, _, _ := g.NodeByName(node1)
	n2, _, _ := g.NodeByName(node2)
	return int(g.underlying.Cost(n1, n2))
}

func (g *Graph[T]) node(label string) int {
	if id, ok := g.nodeIds[label]; ok {
		return id
	} else {
		newId := len(g.nodeNames)
		var none T
		g.nodeNames = append(g.nodeNames, label)
		g.nodeContext = append(g.nodeContext, none)
		g.nodeIds[label] = newId
		return newId
	}
}

func (g *Graph[T]) Print() {
	//fmt.Printf("Node IDs: %v\n", g.nodeIds)
	fmt.Printf("Names: %v\n", g.nodeNames)
	//fmt.Println(g.nodeNames)
	//fmt.Println(g.nodeContext)
	fmt.Println(g.underlying)
}

// Head returns the top of a directed acyclic graph
func (g *Graph[T]) Head() (int, *string, *T) {
	return g.NodeById(g.TopSort()[0])
}

// TopSort returns a list of node IDs in topological order
// the first result is the head of the DAG
func (g *Graph[T]) TopSort() []int {
	topo, ok := graph.TopSort(g.Underlying())
	if !ok {
		panic("no topological sort available; is this a DAG?")
	}

	// the original TopSort() has a weird bug where it returns all the
	// potential nodes that were never actually allocated,
	// so we need to filter those out
	topo = f8l.Filter(topo, func(v int) bool { return v < len(g.nodeNames) })
	return topo
}

// TODO: is this pre-order or in-order?
// TODO: does it also make sense to implement this as an iterator?
func (g *Graph[T]) DFS(origin int, fn func(v int, label string, ctx *T, path []int)) {
	s := stack.NewStack[int]()
	s.Push(origin)
	parents := make([]int, g.Underlying().Order())
	parents[origin] = -1
	path := func(v int) []int {
		pathStack := stack.NewStack[int]()
		for {
			parent := parents[v]
			if parent == -1 {
				return pathStack.Iterator().List()
			} else {
				pathStack.Push(parent)
				v = parent
			}
		}
	}
	for !s.Empty() {
		visitNodeId := s.MustPop()
		_, label, ctx := g.NodeById(visitNodeId)
		fn(visitNodeId, *label, ctx, path(visitNodeId))
		g.Underlying().Visit(visitNodeId, func(w int, c int64) bool {
			s.Push(w)
			parents[w] = visitNodeId
			return false
		})
	}

}

// the underlying does not support disconnected nodes
// so there's no point implementing this
//func (g *Graph[T]) AddNode(Label string) {
//
//}
