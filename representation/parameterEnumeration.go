package representation

type ParameterEnumeration interface {
	Get(uint64) Parameter
	Length() uint64
}
