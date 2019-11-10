package roman

import (
	"strings"
)

var numerals = map[int]string{
	1:  "I",
	5:  "V",
	10: "X",
}

func Roman(number int) string {
	if number >= 0 && number <= 3 {
		return strings.Repeat("I", number)
	}

	if number == 4 {
		return "IV"
	}

	if number == 5 {
		return "V"
	}

	if number >= 6 && number <= 8 {
		return "V" + strings.Repeat("I", number-5)
	}

	if number == 9 {
		return "IX"
	}

	if number == 10 {
		return "X"
	}

	if number >= 11 && number <= 13 {
		return "X" + strings.Repeat("I", number-10)
	}

	if number == 15 {
		return "XV"
	}

	return ""
}
