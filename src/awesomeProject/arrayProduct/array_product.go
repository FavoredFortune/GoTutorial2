package arrayProduct


// given an array of integers return an array of products of all elements in that array,
// excluding the value at that index

// input [2,4,3,5] output // [60,40, 20, 24]

//if array is empty return empty array
// numbers can be negative or positive
// product will never hit integer overflow (max/min)

func arrayProduct(input []int) []int{
	if len(input) == 0 {
		return []int(nil)
	}
	var result []int
	temp := 1
	var zeroTracker int
	
	for _, v := range input {
		if v == 0 {
			zeroTracker ++
		}
		if v != 0 {
			temp = temp * v
		}
	}
	
	var product int
	for _, v := range input {
		if zeroTracker > 1{
			result = append(result, 0)
		} else if v == 0 && zeroTracker == 1 {
			result = append(result, temp)
		} else if v != 0 && zeroTracker >= 1{
			result = append(result, 0)
		} else {
			product = temp / v
			result = append(result, product)
		}
	}
	return result
}
