package arrayandslices

func Sum(numbers []int) int {

	var sum int = 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}