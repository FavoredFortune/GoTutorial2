package nextGreatestLexicographical

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNL(t *testing.T){
	t.Run("it finds next largest lexicograpical number", func(t *testing.T) {
		input := []int{2,1,8,7,6,5}
		
		result, err := findNL(input)
		
		assert.NoError(t, err)
		assert.NotNil(t,result)
		assert.Equal(t, 251678, result)
	})
	
}
func TestMakeValue(t *testing.T){
	t.Run("it turns arrays into numbers", func(t *testing.T) {
		input := []int{2,3,5,6,6}
		
		result := makeValue(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 23566, result)
	})
	t.Run("it handles small numbers", func(t *testing.T) {
		input := []int{2,3}
		
		result := makeValue(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 23, result)
	})
	t.Run("it handles 0", func(t *testing.T) {
		input := []int{2,0}
		
		result := makeValue(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 20, result)
	})
	t.Run("it handles 0 in several places", func(t *testing.T) {
		input := []int{2,0,3,2,0,7}
		
		result := makeValue(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 203207, result)
	})
}

func TestSquare(t *testing.T){
	t.Run("it squares values of ten squared for creating numbers based on array length", func(t *testing.T) {
		input := 4
		
		result := square(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 1000,result)
	})
}

func TestSwap(t *testing.T){
	t.Run("it swaps larger and smaller numbers correctly", func(t *testing.T) {
		input := []int{3,4}
		
		result := swap(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, []int{4,3}, result)
	})
	t.Run("how does it handle longer arrays?", func(t *testing.T) {
		input := []int{2,3,4}
		
		result := swap(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, []int{4, 3, 2}, result)
	})
}

