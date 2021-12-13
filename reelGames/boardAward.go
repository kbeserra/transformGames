package reelGames

import (
	"fmt"
	"github.com/kbeserra/transformGames/representation"
	"strings"
)

type BoardAward struct {
	name   string
	cells  [][]int
	awards []representation.Award
}

func (ba BoardAward) Value() float64 {
	return representation.AwardValueSum(ba.awards)
}

func (ba BoardAward) String() string {
	var sb strings.Builder
	sb.WriteString(ba.name)
	if ba.cells != nil {
		sb.WriteString(" cells: [")
		for i, col := range ba.cells {
			sb.WriteString("[")
			for j, r := range col {
				sb.WriteString(fmt.Sprintf("%d", r))
				if j+1 < len(col) {
					sb.WriteString(", ")
				}
			}
			sb.WriteString("]")
			if i+1 < len(ba.cells) {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]")
	}
	return sb.String()
}

func (ba BoardAward) SubAwards() []representation.Award {
	return ba.awards
}

func (ba BoardAward) Copy() representation.Award {
	awards := make([]representation.Award, len(ba.awards))
	for i, a := range ba.awards {
		awards[i] = a.Copy()
	}
	cells := make([][]int, len(ba.cells))
	for i, col := range ba.cells {
		cells[i] = make([]int, len(col))
		copy(cells[i], col)
	}

	return &BoardAward{
		name:   ba.name,
		cells:  cells,
		awards: awards,
	}
}
