package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
