package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	args := struct {
		s     string
		count int
	}{"a", 8}
	want := "aaaaaaaa"

	if got := Repeat(args.s, args.count); got != want {
		t.Errorf("Repeat(%q, %d) = %q, want %q", args.s, args.count, got, want)
	}
}

func TestRepeatBuf(t *testing.T) {
	args := struct {
		s     string
		count int
	}{"a", 8}
	want := "aaaaaaaa"

	if got := Repeat2(args.s, args.count); got != want {
		t.Errorf("Repeat(%q, %d) = %q, want %q", args.s, args.count, got, want)
	}
}

// Initial implementation with simple string addition (r += s)
// 	BenchmarkRepeat-8   	 2539945	       466 ns/op	      64 B/op	       9 allocs/op
// Second implementation with slice allocated to wanted capacity
//	BenchmarkRepeat-8   	 3752137	       320 ns/op	     176 B/op	       2 allocs/op
// Third implementation with strings.Builder
// 	BenchmarkRepeat-8   	10664800	       105 ns/op	      16 B/op	       1 allocs/op
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a", 100)
	}
}

func ExampleRepeat() {
	s := Repeat("a", 5)
	fmt.Println(s)
	//Output: aaaaa
}
