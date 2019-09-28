package iteration

import "strings"

func Repeat(s string, count int) string {
	b := strings.Builder{}
	b.Grow(len(s) * count)

	for i := 0; i < count; i++ {
		b.WriteString(s)
	}

	return b.String()
}
