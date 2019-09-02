package findUnique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUnique(t *testing.T){
	t.Run("it works with larger array", func(t *testing.T) {
		input := []int{1, 1, 3, 7, 2, 7, 2, 4, 9, 4, 9}
		
		result := findUnique(input)
		assert.NotNil(t,result)
		assert.Equal(t, 3, result)
	})
	
	t.Run("it works with tiny array", func(t *testing.T) {
		input := []int{1, 1, 3}
		
		result := findUnique(input)
		assert.NotNil(t,result)
		assert.Equal(t, 3, result)
	})
	
	t.Run("it works with 0s", func(t *testing.T) {
		input :=[]int {0, 0, 7, 7, 1}
		
		result := findUnique(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 1, result)
	})
	
	t.Run("it works with negative numbers", func(t *testing.T) {
		input := []int{-3, -1, 32, 32, -7, -3, -1}
		
		result := findUnique(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, -7, result)
	})
	
}
