package roman

import (
	"fmt"
	"testing"
)

func TestRoman(t *testing.T) {
	cases := []struct {
		number int
		want   string
	}{
		{1, "I"},
		{2, "II"},
		{4, "IV"},
		{5, "V"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{13, "XIII"},
		{15, "XV"},
	}
	for _, test := range cases {
		name := fmt.Sprintf("%d returns %s", test.number, test.want)
		t.Run(name, func(t *testing.T) {
			if got := Roman(test.number); got != test.want {
				t.Errorf("Roman(%d) = %q, want %q", test.number, got, test.want)
			}
		})
	}
}
