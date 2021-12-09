package reelGames

type symbolSequence struct {
	name       string
	offSet     int
	components []map[string]struct{}
}

func (s symbolSequence) matchesLine(line []string) bool {
	if len(line) < s.offSet+len(s.components) {
		return false
	}

	for i, sym := range line[s.offSet:] {
		if len(s.components) <= i {
			break
		}
		if _, in := s.components[i][sym]; !in {
			return false
		}
	}

	return true
}

func (s symbolSequence) numWays(cols [][]string) int {
	if len(cols) < s.offSet+len(s.components) {
		return 0
	}

	count, col_count := 1, 0
	for i, col := range cols[s.offSet:] {
		if len(s.components) <= i {
			break
		}
		col_count = 0
		for _, sym := range col {
			if _, in := s.components[i][sym]; in {
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

func (s symbolSequence) subSetEq(t symbolSequence) bool {

	// t must be "shorter" than s. That is, t makes few restrictions.
	if t.offSet < s.offSet ||
		s.offSet+len(s.components) < t.offSet+len(t.components) {
		return false
	}

	// each of t's components must make fewer restrictions.
	for i, a := range t.components {
		// a is representing the (t.offSet + i)th compnent.
		// we must compare that to the corresponding compnent of s.
		b := s.components[t.offSet-s.offSet+i]
		// need b to be a subset of a
		for sym := range b {
			if _, in := a[sym]; !in {
				return false
			}
		}
	}

	return true
}
