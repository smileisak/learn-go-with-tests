package arrays

// Sum calculates the sum of an array
func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculate the sum of slices and return the result in a slice
func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	// return sums
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails calculates the totals of the "tails" of each slice.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}

	}
	return sums
}
