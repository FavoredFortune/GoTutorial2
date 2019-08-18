package reverseInt

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestReverseInt(t *testing.T){
	t.Run("it reverses a positive integer", func(t *testing.T) {
		testInt := 4357
		
		result := reverse(testInt)
		
		assert.NotNil(t, result)
		assert.Equal(t, 7534, result)
	})
	t.Run("it reverses a negative integer", func(t *testing.T) {
		negInt := -10934
		
		result := reverse(negInt)
		
		assert.NotNil(t, result)
		assert.Equal(t, -43901, result)
	})
	t.Run("it returns zero is number entered is larger than max int", func(t *testing.T) {
		test := math.MaxInt32
		
		result := reverse(test)
		
		assert.NotNil(t, result)
		assert.Equal(t,0, result)
	})
}

func TestReverseTwo(t *testing.T){
	t.Run("it reverses a positive integer", func(t *testing.T) {
		testInt := 4357
		
		result := reverseTwo(testInt)
		
		assert.NotNil(t, result)
		assert.Equal(t, 7534, result)
	})
	t.Run("it reverses a negative integer", func(t *testing.T) {
		negInt := -10934
		
		result := reverseTwo(negInt)
		
		assert.NotNil(t, result)
		assert.Equal(t, -43901, result)
	})
	t.Run("it returns zero is number entered is larger than max int", func(t *testing.T) {
		test := math.MaxInt32
		
		result := reverseTwo(test)
		
		assert.NotNil(t, result)
		assert.Equal(t,0, result)
	})
}
