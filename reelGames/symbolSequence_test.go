package reelGames

import "testing"

func TestMatchesLine(t *testing.T) {
	s := SymbolSequence{
		Name:   "testSequence",
		OffSet: 1,
		Components: []map[string]struct{}{
			{"A": struct{}{}, "C": struct{}{}},
			{"A": struct{}{}},
		},
	}

	lines := []struct {
		line  []string
		match bool
		msg   string
	}{
		{line: []string{"A", "A"},
			match: false,
			msg:   "line is too short to match"},
		{line: []string{"D", "B", "A"},
			match: false,
			msg:   "symbol in 2nd component is not in sequence component"},
		{line: []string{"D", "B", "A", "A", "A"},
			match: false,
			msg:   "symbol in 2nd component is not in sequence component"},
		{line: []string{"D", "A", "B"},
			match: false,
			msg:   "symbol in 3rd component is not in sequence component"},
		{line: []string{"D", "A", "B", "B", "D"},
			match: false,
			msg:   "symbol in 3rd component is not in sequence component"},
		{line: []string{"D", "A", "A"},
			match: true,
			msg:   "should match"},
		{line: []string{"D", "A", "A", "D", "D"},
			match: true,
			msg:   "should match"},
		{line: []string{"D", "C", "A"},
			match: true,
			msg:   "should match"},
		{line: []string{"D", "C", "A", "D", "A"},
			match: true,
			msg:   "should match"},
	}

	for _, test := range lines {
		got := s.matchesLine(test.line)
		if got != test.match {
			t.Errorf("line match with %v is %v; expected %v. As, %s", test.line, got, test.match, test.msg)
		}
	}
}

func TestNumWays(t *testing.T) {
	s := SymbolSequence{
		Name:   "testSequence",
		OffSet: 1,
		Components: []map[string]struct{}{
			{"A": struct{}{}, "C": struct{}{}},
			{"A": struct{}{}},
		},
	}

	cols := []struct {
		cols [][]string
		ways int
		msg  string
	}{
		{cols: [][]string{{"A"}, {"A"}},
			ways: 0,
			msg:  "cols is too short to match"},
		{cols: [][]string{{"D"}, {"B", "B"}, {"A"}},
			ways: 0,
			msg:  "symbol in 2nd component has no symbols in the corresponding component"},
		{cols: [][]string{{"D"}, {"B", "B"}, {"A"}, {"A"}, {"A"}},
			ways: 0,
			msg:  "symbol in 2nd component has no symbols in the corresponding component"},
		{cols: [][]string{{"D"}, {"A"}, {"B"}},
			ways: 0,
			msg:  "symbol in 3rd component has no symbols in the corresponding component"},
		{cols: [][]string{{"D"}, {"A"}, {"B"}, {"B"}, {"D"}},
			ways: 0,
			msg:  "symbol in 3rd component has no symbols in the corresponding component"},
		{cols: [][]string{{"D"}, {"A"}, {"A"}},
			ways: 1,
			msg:  "1 way"},
		{cols: [][]string{{"D"}, {"A"}, {"A", "A"}, {"D"}, {"D"}},
			ways: 2,
			msg:  "2 ways"},
		{cols: [][]string{{"D"}, {"A", "C"}, {"A", "A", "B"}},
			ways: 4,
			msg:  "4 ways"},
		{cols: [][]string{{"D"}, {"A", "C"}, {"A", "A", "B"}, {"D"}, {"D"}},
			ways: 4,
			msg:  "4 ways"},
	}

	for _, test := range cols {
		got := s.numWays(test.cols)
		if got != test.ways {
			t.Errorf("cols match with %v is %v; expected %v. As, %s", test.cols, got, test.ways, test.msg)
		}
	}
}

func TestSubSetEq(t *testing.T) {
	s := SymbolSequence{
		Name:   "testSequence",
		OffSet: 1,
		Components: []map[string]struct{}{
			{"A": struct{}{}, "C": struct{}{}},
			{"A": struct{}{}},
		},
	}

	tests := []struct {
		seq        SymbolSequence
		isSubSetEq bool
		msg        string
	}{
		{seq: s,
			isSubSetEq: true,
			msg:        "the two are equal",
		},
		{seq: SymbolSequence{
			OffSet: 1,
			Components: []map[string]struct{}{
				{"A": struct{}{}, "C": struct{}{}},
				{"A": struct{}{}, "B": struct{}{}},
			},
		},
			isSubSetEq: true,
			msg:        "last component allows more options",
		},
		{seq: SymbolSequence{
			OffSet: 1,
			Components: []map[string]struct{}{
				{"A": struct{}{}, "C": struct{}{}},
				{"B": struct{}{}},
			},
		},
			isSubSetEq: false,
			msg:        "last component is disjoint from corresponding component",
		},
		{seq: SymbolSequence{
			OffSet: 0,
			Components: []map[string]struct{}{
				{"A": struct{}{}},
				{"A": struct{}{}, "C": struct{}{}},
				{"A": struct{}{}},
			},
		},
			isSubSetEq: false,
			msg:        "offset inccures more restrictions",
		},

		{seq: SymbolSequence{
			OffSet: 1,
			Components: []map[string]struct{}{
				{"A": struct{}{}, "C": struct{}{}},
				{"A": struct{}{}},
				{"A": struct{}{}},
			},
		},
			isSubSetEq: false,
			msg:        "extra component inccures more restrictions",
		},
	}

	for _, test := range tests {
		got := s.subSetEq(test.seq)
		if got != test.isSubSetEq {
			t.Errorf("is subset got %v; expected %v. As, %s", test.isSubSetEq, got, test.msg)
		}
	}
}
