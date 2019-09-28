package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"add two numbers",
			[]int{2, 2},
			4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args...); got != tt.want {
				t.Errorf("Add(%d) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 2, 3)
	fmt.Println(sum)
	//Output: 6
}
