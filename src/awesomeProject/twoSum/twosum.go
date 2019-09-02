package twoSum

// solution in March 2019
// 40 ms	2.9 MB
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
// 36 ms	3 MB
func TwoSum(nums []int, target int)[]int{
	var result []int
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

// so simple brute force way
// take each element and add it to the next element and test the value against the target
// if the resulting value  equals target return those two indicies in a new slice
// NEED TWO TRACKERS


func twoSum2(nums []int, target int) []int {
	result :=[]int{}

	for i := 0; i<len(nums)-1 ; i ++ {

		for j:= i+1 ; j< len(nums); j++ {
			
			if target == nums[i] + nums[j]{
				result = append(result, i, j)
				return result
			}
		}
		
	}
	return result
}