package arrays

// Receiving array of integers and returns those sum
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Receiving multiple arrays of integers and returning a slice of sums for each array
func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] += Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		if len(numbers) > 0 {
			sums[i] = Sum(numbers[1:])
		} else {
			sums[i] = 0
		}
	}
	return sums
}
