package arrays

// Sum return the sum of
// integers contained in an array
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll return an array of integers sum
// from anothers arrays
func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}
