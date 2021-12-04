package main

import (
	"fmt"

	"github.com/kbeserra/transformGames/reelGames"
	"github.com/kbeserra/transformGames/representation"
)

func main() {

	former := &reelGames.BoardFromReelsMorphism{
		Name:    "former",
		Heights: []uint{3, 3, 3},
		Reels: [][]string{
			{"A", "C", "B", "C", "C", "B"},
			{"A", "C", "B", "C", "C", "B"},
			{"A", "C", "B", "C", "C", "B"},
		},
	}

	rowRemover := &reelGames.RemoveRowMorphism{
		Name: "remover",
		Row:  0,
	}

	game := representation.ConcatenationMorphism{
		Name:         "game",
		Morphisms:    []representation.Morphism{former, rowRemover},
		Expand:       true,
		IgnoreStates: false,
		IgnoreAwards: false,
	}

	params, err := game.EnumerateParameters()
	if err != nil {
		panic(err)
	}

	size := params.EnumerationCardinality()
	weight := params.WeightedEnumerationCardinality()
	fmt.Println(size, weight)
	for i := uint64(0); i < weight; i++ {
		p := params.WeightedEnumeration(i)
		o, err := game.Apply(nil, p)
		if err != nil {
			panic(err)
		}
		fmt.Println("------------------------")
		representation.PrintOutcome(o)
	}

	// idMorphism := &representation.IdentityMorphism{
	// 	Name: "id",
	// }
	//
	// parameterization := &representation.ConcatenationParameterization{
	// 	Parameterizations: []representation.Parameterization{
	// 		&representation.IntegerIntervalParameterization{
	// 			LowerBound: 0,
	// 			UpperBound: 3,
	// 			Weights:    []uint64{3, 2, 1},
	// 		},
	// 		&representation.IntegerIntervalParameterization{
	// 			LowerBound: 5,
	// 			UpperBound: 8,
	// 			Weights:    []uint64{2, 2, 2},
	// 		},
	// 	},
	// }
	//
	// parameterization.Parameterizations[0].AssignMorphism(idMorphism)
	//
	// frozenParameters := map[string]representation.Parameterization{
	// 	idMorphism.Name: &representation.ConstantParameterization{
	// 		C: 0,
	// 	},
	// }
	//
	// frozenPerameterization := parameterization.Freeze(frozenParameters)
	//
	// size := frozenPerameterization.EnumerationCardinality()
	// weight := frozenPerameterization.WeightedEnumerationCardinality()
	// fmt.Println(size, weight)
	// for i := uint64(0); i < weight; i++ {
	// 	p := frozenPerameterization.WeightedEnumeration(i).Value()
	// 	fmt.Println(p)
	// }
	//
	//
}
