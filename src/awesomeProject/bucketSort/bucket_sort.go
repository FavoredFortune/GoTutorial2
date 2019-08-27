package bucketSort

import "fmt"

//https://coderpad.io/7G3HKKZC

func Sort(nums []int) ([]int, error){
	if len(nums) <= 0{
		return nil, fmt.Errorf("input was empty %v", nums)
	}
	// making the counters bucket
	counters := make([]int, 100)
	fmt.Println("SHOW ME COUNTERS \n",  counters)
	
	//count
	for _,v := range nums{
		counters[v] = counters[v] + 1
	}
	fmt.Println("AFTER COUNTERS IS USED \n", counters)
	
	var result []int
	//result
	for i , v := range counters{
		if v != 0 {
			//hacky solution to insist on appends by number of instances in V (i.e.
			// append "v" times)
			//if v == 1 {
			//	result = append(result, i)
			//}
			//if v == 2 {
			//	result = append(result, i)
			//	result = append(result, i)
			//}
			for v > 0 {
				result = append(result, i)
				v --
			}
		}
	}
	
	return result, nil
}

