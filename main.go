package main

import (
	"fmt"

	"github.com/kbeserra/transformGames/representation"
)

func main() {

	idMorphism := &representation.IdentityMorphism{
		Name: "id",
	}

	parameterization := &representation.ConcatenationParameterization{
		Parameterizations: []representation.Parameterization{
			&representation.IntegerIntervalParameterization{
				LowerBound: 0,
				UpperBound: 3,
				Weights:    []uint64{3, 2, 1},
			},
			&representation.IntegerIntervalParameterization{
				LowerBound: 5,
				UpperBound: 8,
				Weights:    []uint64{2, 2, 2},
			},
		},
	}

	parameterization.Parameterizations[0].AssignMorphism(idMorphism)

	frozenParameters := map[string]representation.Parameterization{
		idMorphism.Name: &representation.ConstantParameterization{
			C: 0,
		},
	}

	frozenPerameterization := parameterization.Freeze(frozenParameters)

	size := frozenPerameterization.EnumerationCardinality()
	weight := frozenPerameterization.WeightedEnumerationCardinality()
	fmt.Println(size, weight)
	for i := uint64(0); i < weight; i++ {
		p := frozenPerameterization.WeightedEnumeration(i).Value()
		fmt.Println(p)
	}
}
