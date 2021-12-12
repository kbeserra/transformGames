package reelGames

type SymbolSequence struct {
	Name       string
	OffSet     int
	Components []map[string]struct{}
}

func (s SymbolSequence) matchesLine(line []string) bool {
	if len(line) < s.OffSet+len(s.Components) {
		return false
	}

	for i, sym := range line[s.OffSet:] {
		if len(s.Components) <= i {
			break
		}
		if _, in := s.Components[i][sym]; !in {
			return false
		}
	}

	return true
}

func (s SymbolSequence) numWays(cols [][]string) int {
	if len(cols) < s.OffSet+len(s.Components) {
		return 0
	}

	count, col_count := 1, 0
	for i, col := range cols[s.OffSet:] {
		if len(s.Components) <= i {
			break
		}
		col_count = 0
		for _, sym := range col {
			if _, in := s.Components[i][sym]; in {
				col_count++
			}
		}
		count *= col_count
		if count == 0 {
			break
		}
	}

	return count
}

func (s SymbolSequence) subSetEq(t SymbolSequence) bool {

	// t must be "shorter" than s. That is, t makes few restrictions.
	if t.OffSet < s.OffSet ||
		s.OffSet+len(s.Components) < t.OffSet+len(t.Components) {
		return false
	}

	// each of t's Components must make fewer restrictions.
	for i, a := range t.Components {
		// a is representing the (t.OffSet + i)th compnent.
		// we must compare that to the corresponding compnent of s.
		b := s.Components[t.OffSet-s.OffSet+i]
		// need b to be a subset of a
		for sym := range b {
			if _, in := a[sym]; !in {
				return false
			}
		}
	}

	return true
}
