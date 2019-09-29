package iteration

import (
	"strings"
	"unsafe"
)

func Repeat(s string, count int) string {
	b := strings.Builder{}
	b.Grow(len(s) * count)

	for i := 0; i < count; i++ {
		b.WriteString(s)
	}

	return b.String()
}

// Alternate implementation based on how strings.Builder works internally.
// It only has a very small performance advantage over the other function,
// and it's mostly here because I find the code interesting.
func Repeat2(s string, count int) string {
	b := make([]byte, len(s), len(s)*count)
	copy(b, s)

	for i := 1; i < count; i++ {
		b = append(b, s...)
	}

	return *(*string)(unsafe.Pointer(&b))
}
