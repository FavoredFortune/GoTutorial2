package findDuplicate

func findDuplicate(nums []int) int {
	// if there must be a duplicate in every array submitted to the function then an array with 2 or less numbers must contain it's own duplicate
	if len(nums) <= 2 {
		return nums[0]
	}
	
	check := map[int]bool{}
	
	for _, i := range nums {
		_, ok := check[i]
		if !ok {
			check[i] = true
		}
		if ok {
			return i
		}
	}
	return 0
}
