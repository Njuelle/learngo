package arrays

// Sum return the sum of five
// integers contained in an array
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

//SumAll return an array of sum of another array
func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}
