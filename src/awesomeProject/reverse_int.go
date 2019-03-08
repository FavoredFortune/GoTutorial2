package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	testIntP := 4345
	testIntN := -123

	fmt.Println("and the answer is: ", reverse(testIntP))
	fmt.Println("and the answer is: ", reverse(testIntN))

	fmt.Println("try reverseTwo method")
	fmt.Println( reverseTwo(testIntP))
	fmt.Println( reverseTwo(testIntN))

}

//original solution couldn't handle negatives
//referenced this for help: https://leetcode.com/problems/reverse-integer/discuss/184798/golang
func reverse(x int) int {

	orgS := ""
	newS := ""

	//convert int to string using "i to a" method built into Go
	if x > 0{
		orgS = strconv.Itoa(x)
	} else {
		//in case X is negative
		orgS = strconv.Itoa(-x)
	}

	fmt.Println(orgS)

	for i := len(orgS)-1; i> -1; i--{
		newS += string(orgS[i])
	}

	fmt.Println("new string", newS)

	//convert string into a 32 bit int using "a to i" method built into Go which requires both an int and err as output
	//err can be ignored by using "_" or you can use if statement to include check for err
	newInt, err := strconv.Atoi(newS)
	if err == nil {
		fmt.Println(newInt)
	}

	//add negative back in at the end if part of original value
	if x < 0{
		newInt = -newInt

		fmt.Println(newInt)
	}

	//wasn't passing leetcode tests, because wasn't accounting for integer overflow. this is the fix
	//check for integer overflows (https://stackoverflow.com/questions/33641717/detect-signed-int-overflow-in-go)
	if newInt > math.MaxInt32 || newInt < math.MinInt32 {
		return 0
	}
	return newInt
}


//alt solution: USE MODULUS! (after reading discussion on leetcode)
func reverseTwo (x int) int {

	sum :=0

	for x !=0 {
		sum = sum*10 +x%10
		x = x/10
	}

	//also needed to pass
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum
}