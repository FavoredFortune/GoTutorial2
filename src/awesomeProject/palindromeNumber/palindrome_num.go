package palindromeNumber

import (
	"strconv"
)

// check for empty return true if empty
// check if negative return false
// turn int into string
// declare empty slice
// put string values into that slice
// use two pointers and iterate through string backwards from end,  vs beginng and if values don't match return false
// else return true

func isPalindrome(x int) bool {
	if x < 0{
		return false
	}
	if x == 0 || (1 <= x && x < 10 ){
		return true
	}
	string := strconv.Itoa(x)
	test := []int{}
	for _, r := range string {
		test = append(test,int(r - '0'))
	}
	for i := len(test)-1;  i > 0; i-- {
		for j :=0 ; j < i; j++{
			if test[j] != test[i]{
				return false
			}
			i--
			if j == i {
				return true
			}
		}
	}
	return true
}
