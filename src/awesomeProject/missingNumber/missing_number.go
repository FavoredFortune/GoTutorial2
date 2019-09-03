package missingNumber

// so knowing the length of the array you know the highest number (length = 3 in first example length - 9 in second example)
// make hash of array values
// iterate from 0 to length and find missing number return missing number

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	check := map[int]int{}
	for _, v := range nums {
		check[v] = 1
	}
	for i := 0; i <= len(nums); i++{
		_, ok := check[i]
		if !ok {
			return i
		}
	}
	return 0
}
