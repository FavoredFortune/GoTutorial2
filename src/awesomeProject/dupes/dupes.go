package dupes

import "fmt"

func main() {

	testInputA := []int{1,1,2,3,4}
	testInputB := []int{2,4,7}
	testInputC := []int{}
	testInputD := []int{-2,4,0,-1,9,-2}


	fmt.Println(containsDuplicate(testInputA))
	fmt.Println(containsDuplicate(testInputB))
	fmt.Println(containsDuplicate(testInputC))
	fmt.Println(containsDuplicate(testInputD))

}
//original leetcode problem: https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	//declare and assign boolean to value of false
	result := false

	//edge case if int array is less than 2 there are no dupes possible
	if len(nums) < 2 {
		return result
	}

	//build a map in go with key of the int in the numbs array and add a boolean value
	check := make(map[int]bool)

	//only for loops in go using range is like using a for each loop in Java
	//iterate through passed in int array to add the input ints from the array to the check map
	for _, v := range nums {
		// if the value exists in the map, there is a dupe, so the result is true
		if check[v]{
			//subsequent change/assignment of the value for the variable called result
			return true
		}

		//else each iteration, when you add the int to the check map, assign the bool value of true to the map with the int[] int value
		check[v] = true
	}
	//return final result
	return result
}