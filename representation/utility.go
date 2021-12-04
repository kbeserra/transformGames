package representation

import "fmt"

func ErrFailedToCastFromTo(from, to interface{}) error {
	return fmt.Errorf("failed to case from %T to %T", from, to)
}
