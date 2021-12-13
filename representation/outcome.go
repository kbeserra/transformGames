package representation

import (
	"fmt"
	"sync"
)

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
