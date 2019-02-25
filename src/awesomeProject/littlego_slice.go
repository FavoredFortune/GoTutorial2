package main

import (
	"fmt"
)

func main() {
	source := []int{1,2,3,4,5}
	source = removeAtIndex(source, 2)
	fmt.Println(source)
}

func removeAtIndex(source []int, index int) []int{

	lastIndex := len(source) -1 //index 4 in this case

	fmt.Println("> last index", source[lastIndex]) //value is 5
	fmt.Println(" value at index to remove", source[index])//value 3 at index 2

	source[index], source[lastIndex] = source[lastIndex], source[index] //value 3 = value 5 and value 5 = value 3 so revised array is [1,2,5,4,3]


	fmt.Println(">revised last index", source[lastIndex])
	fmt.Println(" >revised value at index to remove", source[index])

	//return from start to index 4 which is now [1,2,5,4]
	return source[:lastIndex]

}