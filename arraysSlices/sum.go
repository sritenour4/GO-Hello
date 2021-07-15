package arraysSlices

// Sum will take an array of numbers and return the total.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll will take a varying number of slices, returning a new slice containing the totals for each slice passed in.
// SumAll is a variadic function that can take a variable number of arguments (using ...).
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails will calculate the totals of the "tails" of each slice. The tail of a collection is all items in the collection except the first one (the "head").
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // slice the slice - take from 1 to the end with numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
