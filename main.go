package main

import (
	"fmt"

	"github.com/kbeserra/transformGames/representation"
)

func main() {

	enumeration := representation.ConcatenationParameterEnumerateion{
		Enumerations: []representation.ParameterEnumeration{
			&representation.IntegerIntervalParameterEnumerateion{
				LowerBound: 0,
				UpperBound: 3,
				Weights:    []uint64{3, 2, 1},
			},
			&representation.IntegerIntervalParameterEnumerateion{
				LowerBound: 5,
				UpperBound: 8,
				Weights:    []uint64{2, 2, 2},
			},
		},
	}

	size := enumeration.Length()
	weight := enumeration.Weight()
	fmt.Println(size, weight)
	for i := uint64(0); i < weight; i++ {
		p := enumeration.GetWeighted(i).Value()
		fmt.Println(p)
	}
}
