package validAnagram


//create map with count of letters of first string
//each time you find a letter in second string delete a count of that letter from the map
//if the map is empty at the end return true because its an anagram
// else return false

//EDGE CASES
// one empty string == false
// both empty strings == true
// if lengths of strings aren't equal automatically false

// solution works for unicode - not just lowercase

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	if len(s) == 0 || len(t) == 0 {
		return false
	}
	// create an empty map (not nil)
	// can be used for ascii or unicode
	compare := map[rune]int{}
	for _, c:= range s {
		compare[c] +=1
	}
	for _, c := range t {
		_,ok := compare[c]
		if ok {
			compare[c]-=1
			if compare[c] <= 0{
				delete(compare, c)
			}
		}
	}
	if len(compare) == 0 {
		return true
	}
	return false
}