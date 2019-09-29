package arrays_and_slices

func Sum(numbers []int) (sum int) {
	// As it turns out, a range loop like this is actually _faster_
	// than the classical for loop.
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAll(slices ...[]int) []int {
	// The book says to rewrite this into creating a slice with an undefined capacity,
	// and using 'append' to add values to it. That way is a lot slower and harder to read,
	// so let's not do that.
	sums := make([]int, len(slices))

	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	sums := make([]int, len(slices))

	for i, numbers := range slices {
		if len(numbers) == 0 {
			sums[i] = 0
			continue
		}

		sums[i] = Sum(numbers[1:])
	}

	return sums
}
