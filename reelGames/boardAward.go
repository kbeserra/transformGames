package reelGames

import (
	"fmt"
	"github.com/kbeserra/transformGames/representation"
	"strings"
)

type BoardAward struct {
	name   string
	cells  [][2]int
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
		for i, c := range ba.cells {
			row, col := c[BoardCellRow], c[BoardCellColumn]
			sb.WriteString(fmt.Sprintf("(%d, %d)", row, col))
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
	cells := make([][2]int, len(ba.cells))
	for i, cell := range ba.cells {
		for j, x := range cell {
			cells[i][j] = x
		}
	}

	return &BoardAward{
		name:   ba.name,
		cells:  cells,
		awards: awards,
	}
}
