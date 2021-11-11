package main

import (
	"fmt"

	"github.com/kbeserra/transformGames/representation"
)

func main() {

	parameterization := representation.ConcatenationParameterization{
		Parameterizations: []representation.Parameterization{
			&representation.IntegerIntervalParameterization{
				LowerBound: 0,
				UpperBound: 3,
				// Weights:    []uint64{3, 2, 1},
			},
			&representation.IntegerIntervalParameterization{
				LowerBound: 5,
				UpperBound: 8,
				// Weights:    []uint64{2, 2, 2},
			},
		},
	}

	size := parameterization.EnumerationCardinality()
	weight := parameterization.WeightedEnumerationCardinality()
	fmt.Println(size, weight)
	for i := uint64(0); i < weight; i++ {
		p := parameterization.WeightedEnumeration(i).Value()
		fmt.Println(p)
	}
}
