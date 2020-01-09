package arrays

// Sum adds up integers in given Slice
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// SumAll adds up integers in an arbitary number of slices
func SumAll(numbersToSum ...[]int) []int {
	l := len(numbersToSum)
	s := make([]int, l)

	for i, v := range numbersToSum {
		s[i] = Sum(v)
	}

	return s
}

// SumAllTails adds up all ints in slices except for the head ints
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, v := range numbersToSum {
		if len(v) > 1 {
			sums = append(sums, Sum(v[1:]))
		} else {
			sums = append(sums, 0)
		}

	}
	return sums
}
