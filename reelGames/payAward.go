package reelGames

import (
	"fmt"
	"github.com/kbeserra/transformGames/representation"
)

type PayOutAward struct {
	Amount float64
}

func (a PayOutAward) Value() float64 {
	return a.Amount
}

func (a PayOutAward) String() string {
	return fmt.Sprintf("%f", a.Amount)
}

func (a PayOutAward) SubAwards() []representation.Award {
	return nil
}

func (a PayOutAward) Copy() representation.Award {
	return PayOutAward{Amount: a.Amount}
}
