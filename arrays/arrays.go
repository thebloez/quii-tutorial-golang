package main

func Sum(ints []int) int {
	sum := 0
	for _, number := range ints {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// TODO v1
	//lenOfNumbers := len(numbersToSum)
	//sums := make([]int, lenOfNumbers)
	//
	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}
	// TODO v2
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	//	TODO v1
	var sums []int
	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
