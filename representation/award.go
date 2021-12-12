package representation

type Award interface {
	Value() float64
	String() string
	SubAwards() []Award
	Copy() Award
}

func AwardValueSum(awards []Award) float64 {
	rtn := 0.
	for _, a := range awards {
		rtn += a.Value() + AwardValueSum(a.SubAwards())
	}
	return rtn
}
