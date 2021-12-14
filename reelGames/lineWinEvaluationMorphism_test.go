package reelGames

import (
	"github.com/kbeserra/transformGames/representation"
	"testing"
)

func createTestingMorphism() LineWinsEvaluationMorphism {

	return LineWinsEvaluationMorphism{
		Name: "some test",
		Lines: [][]int{
			{0, 0, 0},
			{1, 1, 1},
			{2, 2, 2},
		},
		PaySequences: []SymbolSequence{
			{
				Name:   "3*C",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"C": {}, "W": {}},
					{"C": {}, "W": {}},
					{"C": {}, "W": {}},
				},
			},
			{
				Name:   "3*D",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"D": {}, "W": {}},
					{"D": {}, "W": {}},
					{"D": {}, "W": {}},
				},
			},
			{
				Name:   "2*C",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"C": {}, "W": {}},
					{"C": {}, "W": {}},
				},
			},
			{
				Name:   "2*A",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"A": {}, "W": {}},
					{"A": {}, "W": {}},
				},
			},
			{
				Name:   "3*A",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"A": {}, "W": {}},
					{"A": {}, "W": {}},
					{"A": {}, "W": {}},
				},
			},
			{
				Name:   "3*B",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"B": {}, "W": {}},
					{"B": {}, "W": {}},
					{"B": {}, "W": {}},
				},
			},

			{
				Name:   "2*B",
				OffSet: 0,
				Components: []map[string]struct{}{
					{"B": {}, "W": {}},
					{"B": {}, "W": {}},
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
}

func TestLineWinsEvaluationMorphismInit(t *testing.T) {
	M := createTestingMorphism()
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

func TestLineWinsEvaluationMorphismApply(t *testing.T) {
	M := createTestingMorphism()
	err := M.Init()
	if err != nil {
		t.Error(err)
	}

	originalBoard := Board{
		Symbols: [][]string{
			{"A", "C", "W"},
			{"A", "C", "W"},
			{"A", "D", "A"},
		},
	}

	o := &representation.Outcome{
		Previous: nil,
		M:        nil,
		State: ReelGameState{
			B: originalBoard,
		},
		Awards: nil,
	}

	p, err := M.Apply(o, &representation.ConstantParameter{C: nil})
	if err != nil {
		t.Error(err)
	}
	newState, ok := p.State.(*ReelGameState)
	if !ok {
		t.Errorf("expected new state of %T; got %T", &ReelGameState{}, p.State)
	}
	newBoard := newState.B

	if originalBoard.String() != newBoard.String() {
		t.Error("expected board not to change; board was altered")
	}

	awards := p.AccumulateAwards(nil)
	awardNames := map[string]bool{
		"3*A cells: [(0, 0), (0, 1), (0, 2)]": false,
		"2*C cells: [(1, 0), (1, 1), (1, 2)]": false,
		"3*A cells: [(2, 0), (2, 1), (2, 2)]": false,
	}

	if len(awards) != len(awardNames) {
		t.Errorf("number of returned awards is not %d, got %d", len(awardNames), len(awards))
	}
	for _, a := range awards {
		if _, in := awardNames[a.String()]; in {
			awardNames[a.String()] = true
		}
	}
	for name, in := range awardNames {
		if !in {
			t.Errorf("award with name %s not found in returned awards", name)
		}
	}
	awardValue := M.PayTable["3*A"][0].Value() + M.PayTable["2*C"][0].Value() + M.PayTable["3*A"][0].Value()
	if awardValue != representation.AwardValueSum(awards) {
		t.Errorf("expected award value of %f; got %f", awardValue, representation.AwardValueSum(awards))
	}

}
