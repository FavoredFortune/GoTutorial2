package calculator

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	
	string := "abs"
	string1 := " 2 + 4 * 3 / 9"
	
	fmt.Println(readInput(string))
	fmt.Println(readInput(string1))
	
}
//inspired by https://leetcode.com/problems/basic-calculator-ii/discuss/146256/golang-solution


func readInput(input string)(int, error){
	if len(input) == 0 {
		return 0 , fmt.Errorf("invalid input")
	}
	
	var buff bytes.Buffer
	
	// trims out white space
	for _, r := range input {
		if unicode.IsNumber(r) || r == '+' || r == '-' || r == '/' || r == '*'{
			buff.WriteRune(r)
		}
	}
	
	if buff.Len() == 0 {
		return 0 , fmt.Errorf("invalid input")
	}
	
	// return trimmed buff to string
	string := buff.String()
	
	length := len(string)
	i := 0
	
	for i < length && unicode.IsNumber(rune(string[i])){
		i++
	}
	current , _ := strconv.Atoi(string[:i])
	
	result := 0
	
	for i < length {
		operator := string[i]
		i++
		
		start := i
		for i < length && unicode.IsNumber(rune(string[i])) {i++}
		a, _ := strconv.Atoi(string[start:i])
		
		if operator == '+'{
			result += current
			current = a
		}
		
		if operator == '-'{
			result += current
			current = -1*a
		}
		if operator == '*'{
			current *= a
		}
		
		if operator == '/'{
			current /= a
		}

	}
	
	return result + current, nil
	
}
