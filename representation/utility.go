package representation

import (
	"fmt"
	"github.com/kbeserra/tjson"
)

func ErrFailedToCastFromTo(from, to interface{}) error {
	return fmt.Errorf("failed to case from %T to %T", from, to)
}

func init() {
	tjson.RegisterType("IdentityMorphism", func() interface{} { return &IdentityMorphism{} })
	tjson.RegisterType("ConcatenationMorphism", func() interface{} { return &ConcatenationMorphism{} })
	tjson.RegisterType("ConditionalMorphism", func() interface{} { return &ConditionalMorphism{} })
	tjson.RegisterType("WhileTransformation", func() interface{} { return &WhileMorphism{} })
}
