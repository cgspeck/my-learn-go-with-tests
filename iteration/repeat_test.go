package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	a := Repeat("a", 6)
	e := "aaaaaa"

	if a != e {
		t.Errorf("expected: %q actual: %q", e, a)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("hello", 3)
	fmt.Println(res)
	// Output: hellohellohello
}