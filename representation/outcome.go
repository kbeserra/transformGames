package representation

import (
	"fmt"
	"sync"
)

type Outcome struct {
	Previous  *Outcome
	M         Morphism
	Parameter interface{}
	State     OutcomeState
	Awards    []Award
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

func (o *Outcome) AccumulateParameters(upperBound *Outcome) []interface{} {
	rtn := make([]interface{}, 0, 1)
	o.ApplyToSegment(upperBound, func(p *Outcome) {
		rtn = append(rtn, p.Parameter)
	})
	return rtn
}

func (o *Outcome) AccumulateStates(upperBound *Outcome) []OutcomeState {
	rtn := make([]OutcomeState, 0, 1)
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

func (p Outcome) Copy(weakerBound *Outcome) *Outcome {

	AwardsCopy := make([]Award, len(p.Awards))
	copy(AwardsCopy, p.Awards)

	if p.Previous != nil && p.Previous != weakerBound {
		return &Outcome{
			Previous:  p.Previous.Copy(weakerBound),
			M:         p.M,
			Parameter: p.Parameter,
			State:     p.State.Copy(),
			Awards:    AwardsCopy,
		}
	} else {
		return &p
	}
}

func MergeOutcomeChannels(oChan []<-chan *Outcome) <-chan *Outcome {
	out := make(chan *Outcome)
	var wg sync.WaitGroup
	wg.Add(len(oChan))
	for _, c := range oChan {
		go func(c <-chan *Outcome) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
	Temporary, will be removed with addition of a more appropriate package for this functionality.
*/
func PrintOutcome(o *Outcome) {
	if o.Previous != nil {
		PrintOutcome(o.Previous)
	}
	if o.M != nil {
		fmt.Printf("M: %s\n", o.M.String())
	}
	var stateString string
	if o.State != nil {
		stateString = o.State.String()
	}
	fmt.Printf("State:\n%s\n", stateString)
	fmt.Printf("Awards:\n")
	for _, a := range o.Awards {
		fmt.Printf("\t%s\n", a.String())
	}
}
