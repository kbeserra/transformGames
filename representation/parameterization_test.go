package representation

import "testing"

func TestConstantParameterization(t *testing.T) {

	a := 0
	c := ConstantParameterization{C: a}

	size := c.EnumerationCardinality()
	weight := c.WeightedEnumerationCardinality()

	if size != 1 {
		t.Errorf("size of constant parameterization is %d; want 1", size)
	}
	if weight != 1 {
		t.Errorf("weight of constant parameterization is %d; want 1", size)
	}

	p, w := c.Enumeration(1)
	if w != 1 {
		t.Errorf("weight of value of constant parameterization is %d; want 1", w)
	}
	v, ok := p.Value().(int)
	if !ok {
		t.Errorf("value of constant is of type is %T; want %T", v, a)
	}
	if v != a {
		t.Errorf("value of constant is %v; want %v", v, a)
	}

	// TODO: test Copy and Freeze

}
