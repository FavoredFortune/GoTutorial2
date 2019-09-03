package missingNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T){
	t.Run("it finds the missing number in a short series array", func(t *testing.T) {
		input := []int{3, 0, 1}
		
		result := missingNumber(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 2, result)
	})
	
	t.Run("it finds the missing number in a long series array", func(t *testing.T) {
		input := []int{3, 0, 1, 2, 6, 5, 7, 4, 9}
		
		result := missingNumber(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 8, result)
	})
	
	t.Run("it returns 0 for an empty array", func(t *testing.T) {
		input := []int{}
		
		result := missingNumber(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 0, result)
	})
	
	t.Run("it returns 0 when missing from an array", func(t *testing.T) {
		input := []int{1, 3, 4, 2}
		
		result := missingNumber(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 0, result)
	})
}
