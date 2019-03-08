package main

import "fmt"

func main() {
	nums := []int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}


func singleNumber(nums []int) int {
	
	if len(nums) < 2 {
		return nums[0]
	}
	
	var seen= map[int]int{}
	
	count := 1
	for i := range nums {
		if _, ok := seen[nums[i]]; !ok {
			
			seen[nums[i]] = count
		} else {
			delete(seen, nums[i])
		}
	}
	 for key :=range seen {
	 	return key
	 }
	return 0
	
}