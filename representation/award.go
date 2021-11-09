package representation

type Award interface {
	Value()
	SubAwards()
}
