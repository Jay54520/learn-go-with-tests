package main

func Sum(numbers []int) int  {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) (sum []int)  {

	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sum = append(sum, Sum(tail))
	}

	return
}