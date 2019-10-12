package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	want := 15

	if got := Sum(numbers); got != want {
		t.Errorf("Sum(%d) = %d, want %d", numbers, got, want)
	}

}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func TestSumAll(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{"sum of two slices returns a slice of two results",
			[][]int{
				{1, 2, 3},
				{1, 2, 3, 4, 5},
			},
			[]int{6, 15}},

		{"passing no arguments returns an empty slice",
			[][]int{},
			[]int{}},

		{"passing an empty slice returns a result of 0",
			[][]int{
				{},
			},
			[]int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll(%d) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}

func BenchmarkSumAll(b *testing.B) {
	slices := [][]int{
		{1, 2, 3},
		{1, 2, 3, 4, 5},
	}
	for i := 0; i < b.N; i++ {
		SumAll(slices...)
	}
}

func TestSumAllTails(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{"summing two slices returns two results",
			[][]int{
				{1, 2, 3},
				{1, 2, 3, 4, 5},
			},
			[]int{5, 14}},

		{"passing no arguments returns an empty slice",
			[][]int{},
			[]int{}},

		{"passing an empty slice returns one result of 0",
			[][]int{
				{},
			},
			[]int{0}},
	}
	for _, tt := range tests {
		if got := SumAllTails(tt.args...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SumAllTails(%d) = %d, want %d", tt.args, got, tt.want)
		}
	}
}
