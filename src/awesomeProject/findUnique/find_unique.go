package findUnique

// Given non-empty array of integers where all integers come in pairs except one,
// find the unique integer and return that value.

func findUnique(input []int) int{
	m := map[int]int {}
	
	for _, v := range input {
		_,ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			delete(m, v)
		}
	}
	var result int
	for i := range m {
		result = i
	}
	return result
}

