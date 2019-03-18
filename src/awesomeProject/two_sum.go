package main

import "fmt"

func main() {
	nums := []int {2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
	
	nums2 := []int {3,2,4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))
	
	nums3 := []int {3,2,3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3))
}

func twoSum(nums []int, target int) []int {
	
	for i := 0; i < len(nums); i ++ {
		for j := i + 1 ; j < len(nums); j++ {
			if target - nums[j] - nums[i] == 0 {
				return []int {i, j}
			}
		}
	}
	return []int{}
}