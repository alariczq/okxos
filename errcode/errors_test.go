package errcode

import (
	"fmt"
	"testing"
)

func TestFromError(t *testing.T) {
	ee := New(123, "test")
	eee := fmt.Errorf("wrap: %w", ee)
	e := FromError(eee)
	if e == nil {
		t.Errorf("expected error to be of type *Error")
	}
}
