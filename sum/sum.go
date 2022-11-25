package sum

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersOfSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersOfSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func tails(numbers []int) int {
	lengthOfNumbers := len(numbers)
	if lengthOfNumbers == 0 {
		return 0
	}
	
	return numbers[len(numbers)-1]
}

func SumAllTails(numberOfSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberOfSum {
		sums = append(sums, tails(numbers))
	}

	return sums
}
