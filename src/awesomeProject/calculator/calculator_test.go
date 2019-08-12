package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadInput(t *testing.T){
	t.Run("it returns a value with positive numbers", func(t *testing.T) {
		string1 := "2 + 3 * 4 / 1"
		
		result, err := readInput(string1)
		
		assert.NoError(t, err)
		assert.Equal(t, 14,result)
	})
	
	t.Run("it returns value when some of the input is negative", func(t *testing.T) {
		string := "-3 * 6"
		
		result, err := readInput(string)
		
		assert.NoError(t, err)
		assert.Equal(t, -18,result)
	})
	
	t.Run("it returns an error and 0 when input is invalid", func(t *testing.T) {
		string := "abs d"
		
		result, err := readInput(string)
		
		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
	
}