package integers

// Add takes integers and returns the sum of them.
func Add(n ...int) (sum int) {
	for _, x := range n {
		sum += x
	}

	return
}
