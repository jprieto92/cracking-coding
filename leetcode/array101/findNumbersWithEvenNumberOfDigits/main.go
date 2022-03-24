package findNumbersWithEvenNumberOfDigits

func findNumbers(nums []int) int {
	var count = 0
	for _, element := range nums {
		var digits = 0
		for element > 0 {
			element /= 10
			digits++
		}
		if digits%2 == 0 {
			count++
		}
	}
	return count
}
