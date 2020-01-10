package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	a := buffer.String()
	e := "Hello, Chris"

	if a != e {
		t.Errorf("expected: %q actual: %q", e, a)
	}
}
