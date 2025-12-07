package grid2

func (g *Grid[T]) Equal(g2 *Grid[T]) bool {
	if !g.bounds.Eq(g2.bounds) {
		return false
	}
	if len(g.storage) != len(g2.storage) {
		return false
	}
	for i := range g.storage {
		if g.storage[i] != g2.storage[i] {
			return false
		}
	}
	return true
}
