package binarySearch

// FOR INTS IN A SLICE Positive & negative numbers

func ReturnNewBinarySort(nums []int, target int) []int{
	//address edge case of a empty/nil slice being passed in
	if nums == nil {
		return []int{target}
	}
	var result []int // make slice to return
	x := FindPosition(nums, target) //find where target belongs in slice
	for i := range nums {
		if i == x {
			result = append(result, target)
		}
		result = append(result, nums[i])
	}
	return result
}


// returns what indices the element should be placed in to keep the slice sorted
func FindPosition(nums []int, target int) int {
	for i := range nums {
		if target == nums[i] {
			return i
		} else if target < nums[i] {
			return i
		}
	}
	return len(nums)
}

func VerifyBinary(nums []int)bool{
	var temp int
	for i := range nums {
		temp = i +1
		if nums[i] > temp {
			return false
		}
	}
	return true
}

// returns the index of the target element in the slice, if available or -1 if not
func BiSearch(nums []int, target int) int {
	for i := range nums {
		if target > nums[i] {
			i ++
		} else if target == nums[i] {
			return i
		} else if target < nums [i]{
			if i > 0 {
				i--
			}
			return -1
		}
	}
	return -1
}
