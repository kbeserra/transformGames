package representation

type Outcome struct {
	Previous *Outcome
	M        Morphism
	State    GameState
	Awards   []Award
}

func (o *Outcome) ApplyToSegment(upperBound *Outcome, f func(*Outcome)) {
	if o.Previous != nil && o.Previous != upperBound {
		o.Previous.ApplyToSegment(upperBound, f)
	}
	f(o)
}

func (o *Outcome) AccumulateMorphisms(upperBound *Outcome) []Morphism {
	rtn := make([]Morphism, 0, 1)
	o.ApplyToSegment(upperBound, func(p *Outcome) {
		rtn = append(rtn, p.M)
	})
	return rtn
}

func (o *Outcome) AccumulateStates(upperBound *Outcome) []GameState {
	rtn := make([]GameState, 0, 1)
	o.ApplyToSegment(upperBound, func(p *Outcome) {
		rtn = append(rtn, p.State)
	})
	return rtn
}

func (o *Outcome) AccumulateAwards(upperBound *Outcome) []Award {
	rtn := make([]Award, 0, 1)
	o.ApplyToSegment(upperBound, func(p *Outcome) {
		rtn = append(rtn, p.Awards...)
	})
	return rtn
}
