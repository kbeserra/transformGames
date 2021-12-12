package reelGames

import (
	"github.com/kbeserra/transformGames/representation"
	"testing"
)

func TestLineWinsEvaluationMorphismInit(t *testing.T) {
	M := LineWinsEvaluationMorphism{
		Name: "some test",
		Lines: [][]uint{
			{0, 0, 0},
			{1, 1, 1},
			{2, 2, 2},
		},
		PaySequences: []SymbolSequence{
			{
				Name:   "3*C",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"C": struct{}{}},
					{"C": struct{}{}},
					{"C": struct{}{}},
				},
			},
			{
				Name:   "3*D",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"D": struct{}{}},
					{"D": struct{}{}},
					{"D": struct{}{}},
				},
			},
			{
				Name:   "2*C",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"C": struct{}{}},
					{"C": struct{}{}},
				},
			},
			{
				Name:   "2*A",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"A": struct{}{}},
					{"A": struct{}{}},
				},
			},
			{
				Name:   "3*A",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"A": struct{}{}},
					{"A": struct{}{}},
					{"A": struct{}{}},
				},
			},

			{
				Name:   "3*B",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"B": struct{}{}},
					{"B": struct{}{}},
					{"B": struct{}{}},
				},
			},

			{
				Name:   "2*B",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"B": struct{}{}},
					{"B": struct{}{}},
				},
			},
		},
		PayTable: map[string][]representation.Award{
			"3*A": {PayOutAward{
				8,
			}},
			"2*A": {PayOutAward{
				4,
			}},
			"3*B": {PayOutAward{
				4,
			}},
			"2*B": {PayOutAward{
				2,
			}},
			"3*C": {PayOutAward{
				2,
			}},
			"2*C": {PayOutAward{
				1,
			}},
		},
	}
	err := M.Init()
	if err != nil {
		t.Error(err)
	}

	names := []string{"3*A",
		"2*A",
		"3*B",
		"3*C",
		"2*B",
		"2*C",
		"3*D"}
	for i, s := range M.PaySequences {
		if names[i] != s.Name {
			t.Errorf("incorrect order. Expected %s; got %s", names[i], s.Name)
		}
	}
}
