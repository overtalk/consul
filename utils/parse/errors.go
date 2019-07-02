package parse

import (
	"fmt"
)

// Err : all
type Err struct {
	method string
	origin interface{}
}

func (e Err) Error() string {
	return fmt.Sprintf("failed to parse [%v1] to %s", e.origin, e.method)
}
