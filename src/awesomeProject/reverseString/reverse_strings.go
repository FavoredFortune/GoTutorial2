package reverseString

import "fmt"

func main() {

	testByteA := []byte{'h','e','l','l','o'}

	fmt.Println(testByteA)
	fmt.Println(reverseString(testByteA))

	testByteB := []byte{'s','o','o','z'}

	fmt.Println(testByteB)
	fmt.Println(reverseString(testByteB))
	
}

//Original problem https://leetcode.com/problems/reverse-string/

func reverseString(s []byte) []byte {

	//declare and assign the variables for going through the array from both ends
	i :=0
	j:= len(s)-1

	//use the go for loop like a while loop to iterate through the array
	for i < j {

		//swap values from either end of the array
		s[i],s[j] = s[j], s[i]

		//increment and decrement from both sides of the array and you'll meet in the middle
		i++
		j--
	}

	//return the array
	return s
}