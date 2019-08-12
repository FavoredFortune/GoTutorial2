package removeDupesSorted

import "fmt"

func main() {
	testArr := []int{1, 1, 2}
	fmt.Println(removeDuplicates(testArr))

	test2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(test2))
}

// solution help from https://leetcode.com/problems/remove-duplicates-from-sorted-array/discuss/247190/Go-100-56ms
func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	for i, j := 0, 1; i < j && j < len(nums); {
		if nums[j] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			j--
		}
		i++
		j++
	}
	return len(nums)
}
