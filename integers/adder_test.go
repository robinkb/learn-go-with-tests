package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	x, y := 1, 2
	want := 3
	if got := Add(x, y); got != want {
		t.Errorf("Add(%d, %d) = %d, want %d", x, y, got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	//Output: 3
}
