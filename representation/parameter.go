package representation

type Parameter interface {
	Value() interface{}
}

/*

 */
type ConstantParameter struct {
	C interface{}
}

func (cp *ConstantParameter) Value() interface{} {
	return cp.C
}

/*

 */
type SequenceParameter struct {
	Sigmas []Parameter
}

func (sp *SequenceParameter) Value() interface{} {
	rtn := make([]interface{}, len(sp.Sigmas))
	for i, s := range sp.Sigmas {
		rtn[i] = s.Value()
	}
	return rtn
}
