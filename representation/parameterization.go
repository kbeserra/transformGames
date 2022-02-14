package representation

type Parameterization interface {
	AssignMorphism(Morphism)
	AssociatedMorphism() Morphism

	Enumeration(uint64) (p Parameter, weight uint64)
	WeightedEnumeration(uint64) Parameter

	// Expected to return in constant time.
	EnumerationCardinality() uint64
	// Expected to return in constant time.
	WeightedEnumerationCardinality() uint64

	Copy() Parameterization
	Freeze(map[string]Parameterization) Parameterization
}

/*

 */
type ParameterizationBase struct {
	M Morphism
}

/*

 */
type ConstantParameterization struct {
	ParameterizationBase
	C interface{}
}

func (cpe *ConstantParameterization) AssignMorphism(M Morphism) {
	cpe.M = M
}

func (cpe *ConstantParameterization) AssociatedMorphism() Morphism {
	return cpe.M
}

func (cpe *ConstantParameterization) Enumeration(n uint64) (p Parameter, weight uint64) {
	return &ConstantParameter{C: cpe.C}, 1
}

func (cpe *ConstantParameterization) WeightedEnumeration(uint64) Parameter {
	return &ConstantParameter{C: cpe.C}
}

func (cpe *ConstantParameterization) EnumerationCardinality() uint64 {
	return 1
}

func (cpe *ConstantParameterization) WeightedEnumerationCardinality() uint64 {
	return 1
}

func (cpe *ConstantParameterization) Copy() Parameterization {
	return &ConstantParameterization{
		ParameterizationBase: ParameterizationBase{
			M: cpe.M,
		},
		C: cpe.C,
	}
}

func (cpe *ConstantParameterization) Freeze(frozen map[string]Parameterization) Parameterization {
	if cpe.AssociatedMorphism() != nil {
		if p, in := frozen[cpe.AssociatedMorphism().String()]; in {
			return p
		}
	}
	return cpe.Copy()
}

/*

 */
type SequenceParameterization struct {
	ParameterizationBase
	Parameterizations []Parameterization

	cached bool
	length uint64
	weight uint64
}

func (sp *SequenceParameterization) AssignMorphism(M Morphism) {
	sp.M = M
}

func (sp *SequenceParameterization) AssociatedMorphism() Morphism {
	return sp.M
}

func (sp *SequenceParameterization) cache() {
	sp.length = 1
	sp.weight = 1
	for _, p := range sp.Parameterizations {
		sp.length *= p.EnumerationCardinality()
		sp.weight *= p.WeightedEnumerationCardinality()
	}
	sp.cached = true
}

func (sp *SequenceParameterization) Enumeration(n uint64) (Parameter, uint64) {
	if !sp.cached {
		sp.cache()
	}
	rtn := make([]Parameter, len(sp.Parameterizations))
	var w, weight uint64 = 0, 0
	for i, p := range sp.Parameterizations {
		l := p.EnumerationCardinality()
		rtn[i], w = p.Enumeration(n % l)
		weight *= w
		n /= l
	}

	return &SequenceParameter{
		Sigmas: rtn,
	}, weight
}

func (sp *SequenceParameterization) WeightedEnumeration(n uint64) Parameter {
	if !sp.cached {
		sp.cache()
	}
	rtn := make([]Parameter, len(sp.Parameterizations))
	for i, p := range sp.Parameterizations {
		l := p.WeightedEnumerationCardinality()
		rtn[i] = p.WeightedEnumeration(n % l)
		n /= l
	}

	return &SequenceParameter{
		Sigmas: rtn,
	}
}

func (sp *SequenceParameterization) EnumerationCardinality() uint64 {
	if !sp.cached {
		sp.cache()
	}
	return sp.length
}

func (sp *SequenceParameterization) WeightedEnumerationCardinality() uint64 {
	if !sp.cached {
		sp.cache()
	}
	return sp.weight
}

func (sp *SequenceParameterization) Copy() Parameterization {
	parameterizations := make([]Parameterization, len(sp.Parameterizations))
	for i, p := range sp.Parameterizations {
		parameterizations[i] = p.Copy()
	}
	return &SequenceParameterization{
		ParameterizationBase: ParameterizationBase{
			M: sp.M,
		},
		Parameterizations: parameterizations,
	}
}

func (sp *SequenceParameterization) Freeze(frozen map[string]Parameterization) Parameterization {
	if sp.AssociatedMorphism() != nil {
		if p, in := frozen[sp.AssociatedMorphism().String()]; in {
			return p
		}
	}
	parameterizations := make([]Parameterization, len(sp.Parameterizations))
	for i, p := range sp.Parameterizations {
		parameterizations[i] = p.Freeze(frozen)
	}
	return &SequenceParameterization{
		ParameterizationBase: ParameterizationBase{
			M: sp.M,
		},
		Parameterizations: parameterizations,
	}
}

/*

 */
type IntegerIntervalParameterization struct {
	ParameterizationBase
	// first coordinate is the lower bound, second is the upper bound.
	// We include the lower bound, exclude the upper.
	LowerBound int
	UpperBound int
	Weights    []uint64

	cached            bool
	cumulativeWeights []uint64
}

func (iipe *IntegerIntervalParameterization) AssignMorphism(M Morphism) {
	iipe.M = M
}

func (iipe *IntegerIntervalParameterization) AssociatedMorphism() Morphism {
	return iipe.M
}

func (iipe *IntegerIntervalParameterization) cache() {
	iipe.cumulativeWeights = make([]uint64, uint64(iipe.UpperBound-iipe.LowerBound))
	var w uint64 = 0
	for i := 0; i < (iipe.UpperBound - iipe.LowerBound); i++ {
		if iipe.Weights != nil {
			w += iipe.Weights[i]
		} else {
			w++
		}
		iipe.cumulativeWeights[i] = w
	}
	iipe.cached = true
}

func (iipe *IntegerIntervalParameterization) Enumeration(n uint64) (Parameter, uint64) {
	if !iipe.cached {
		iipe.cache()
	}
	n %= uint64(iipe.UpperBound - iipe.LowerBound)
	var weight uint64 = 1
	if iipe.Weights != nil {
		weight = uint64(iipe.Weights[n])
	}

	return &ConstantParameter{
		C: iipe.LowerBound + int(n),
	}, weight
}

func (iipe *IntegerIntervalParameterization) WeightedEnumeration(n uint64) Parameter {
	if !iipe.cached {
		iipe.cache()
	}
	var i uint64 = 0
	for ; iipe.cumulativeWeights[i] <= n; i++ {
	}
	return &ConstantParameter{
		C: iipe.LowerBound + int(i),
	}
}

func (iipe *IntegerIntervalParameterization) EnumerationCardinality() uint64 {
	if !iipe.cached {
		iipe.cache()
	}
	return uint64(iipe.UpperBound - iipe.LowerBound)
}

func (iipe *IntegerIntervalParameterization) WeightedEnumerationCardinality() uint64 {
	if !iipe.cached {
		iipe.cache()
	}
	return iipe.cumulativeWeights[len(iipe.cumulativeWeights)-1]
}

func (iipe *IntegerIntervalParameterization) Copy() Parameterization {
	var weights []uint64
	if iipe.Weights != nil {
		weights = make([]uint64, len(iipe.Weights))
		copy(weights, iipe.Weights)
	}
	return &IntegerIntervalParameterization{
		ParameterizationBase: ParameterizationBase{
			M: iipe.M,
		},
		LowerBound: iipe.LowerBound,
		UpperBound: iipe.UpperBound,
		Weights:    weights,
	}
}

func (iipe *IntegerIntervalParameterization) Freeze(frozen map[string]Parameterization) Parameterization {
	if iipe.AssociatedMorphism() != nil {
		if p, in := frozen[iipe.AssociatedMorphism().String()]; in {
			return p
		}
	}
	return iipe.Copy()
}
