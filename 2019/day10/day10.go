package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2019, 10)
	aoc.Test(run, "sample1.txt", "8", "")
	aoc.Test(run, "sample2.txt", "33", "")
	aoc.Test(run, "sample3.txt", "35", "")
	aoc.Test(run, "sample4.txt", "41", "")
	aoc.Test(run, "sample5.txt", "210", "")
	aoc.Test(run, "input.txt", "247", "")
	aoc.Run(run)
}

type observation struct {
	base grid.Cell
	vec  []int
	obs  grid.Cell
}

func rayVectors(scale int) []grid.Cell {
	p := grid.NewGrid[uint8](scale, scale)
	const Viewpoint = 1
	const Obscured = 2
	for out := 1; out < scale; out++ {
		edge := iter.Chain(
			grid.Line(grid.Cell{0, out}, grid.Cell{out, out}),
			grid.Line(grid.Cell{out, out}, grid.Cell{out, 0}),
		)
		for edge.Next() {
			cell := edge.Value()
			if p.Get(cell) == Obscured {
				continue
			} else {
				p.Set(cell, Viewpoint)
			}
			ray := p.Ray(grid.Cell{0, 0}, cell)
			_ = ray.Skip(1)
			for ray.Next() {
				p.Set(ray.Value(), Obscured)
			}
		}
	}
	//p.Print()
	// TODO: this should not result in a panic:
	//g := grid.NewGrid[bool](0, 0)
	//g.Grow(scale, false)
	g := grid.NewGridFromSlice[uint8](grid.SliceEnclosing(grid.Cell{scale, scale}, grid.Cell{-scale, -scale}))

	// copy the hardcoded values to the other three quadrants
	// so that all four cardinal directions are totally covered
	r := iter.Range(0, scale).List()
	for _, w := range iter.ProductV(r, r).List() {
		v := p.Get([]int{w[0], w[1]})
		g.Set([]int{w[0], w[1]}, v)
		g.Set([]int{w[0], -w[1]}, v)
		g.Set([]int{-w[0], w[1]}, v)
		g.Set([]int{-w[0], -w[1]}, v)
	}

	//g2 := grid.TransformGrid[int, string](g, func(i int) string {
	//	return map[int]string{0: " ", 1: "*"}[i]
	//})
	//g2.Set(grid.Cell{0, 0}, "O")
	//g.Print()

	ret := make([]grid.Cell, 0)
	for ac := g.Filter(func(c grid.Cell, v uint8) bool { return v == Viewpoint }); ac.Next(); {
		ret = append(ret, ac.Value())
	}

	return ret
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridRunes()

	rayVecList := rayVectors(g.Width())

	var bestCount = 0
	var bestList []observation
	for originIter := g.Cells(); originIter.Next(); {
		// can only sit on asteroids
		origin := originIter.Value()
		if g.Get(origin) != "#" {
			continue
		}
		detectedAsteroids := 0
		findList := make([]observation, 0)
		for _, dir := range rayVecList {
			for cellIter := g.Ray(origin, dir); cellIter.Next(); {
				cell := cellIter.Value()
				val := g.Get(cell)
				if val == "#" {
					findList = append(findList, observation{
						base: util.Copy(origin),
						vec:  util.Copy(dir),
						obs:  util.Copy(cell),
					})
					detectedAsteroids++
					break
				}
			}
		}
		if detectedAsteroids > bestCount {
			bestCount = detectedAsteroids
			bestList = findList
		}
	}

	_ = bestList

	*p1 = strconv.Itoa(bestCount)
	*p2 = strconv.Itoa(0)
}
