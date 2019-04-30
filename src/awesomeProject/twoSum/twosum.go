package twoSum

// solution in March 2019
//func twoSum(nums []int, target int) []int {
//
//	for i := 0; i < len(nums); i ++ {
//		for j := i + 1 ; j < len(nums); j++ {
//			if target - nums[j] - nums[i] == 0 {
//				return []int {i, j}
//			}
//		}
//	}
//	return []int{}
//}

// solution on 30 April 2019
func TwoSum(nums []int, target int)[]int{
	result := []int{}
	if len(nums) < 1{
		return nums
	}
	
	for tracker1 := 0 ; tracker1 < len(nums)-1; tracker1++ {
		for tracker2 := tracker1 +1; tracker2 < len(nums); tracker2++ {
			if target == nums[tracker1] + nums[tracker2] {
				result = append(result, tracker1, tracker2)
				return result
			}
		}
	}
	return result
}