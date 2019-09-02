package nextGreatestLexicographical

import (
	"fmt"
)

//Given a number n, find the smallest number that has same set of digits as n and is greater than n.
// If n is the greatest possible number with its set of digits, then return error "not possible".

//For simplicity of implementation, we have considered input number as an array.

func findNL(lexi []int) ([]int, error){
	if len(lexi) == 0 {
		return nil, fmt.Errorf("input is empty")
	}
	var result []int
	var temp []int
	value := makeValue(lexi)
	for i := range lexi {
		temp = swap(lexi[len(lexi)-i -2:])
		 // appending part isn't correct
		for _, v := range temp{
		 	result = append(result, v)
		 }
		compareResult := makeValue(result)
		fmt.Println(compareResult)
		fmt.Println(value)
		if compareResult > value {
			return result, nil
		}
	}
	return nil, fmt.Errorf("greater number not possible")
}


func makeValue (slice []int)int {
	var value int
	for i, v := range slice {
		place := square(len(slice)-i)
		if place == 0 {
			place = 1
		}
		value += v * place
	}
	return value
}

func square(number int) int {
	result := 1
	for number - 1 > 0 {
		result = result * 10
		number --
	}
	return result
}

func swap(sub []int)[]int{
	var temp int
	for i := range sub {
		if sub[i] < sub[len(sub)-1] {
			temp = sub[i]
			sub[i] = sub[len(sub)-1]
			sub[len(sub)-1] = temp
		}
	}
		return sub
}