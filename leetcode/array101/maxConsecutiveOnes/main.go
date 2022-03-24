package maxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	var max = 0
	var count = 0
	for _, element := range nums {
		if element == 0 {
			if count > max {
				max = count
			}
			count = 0
			continue
		}
		count++
	}
	if count > max {
		max = count
	}
	return max
}
