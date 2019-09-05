package palindromeNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromeNum(t *testing.T)  {
	t.Run("testing large even array", func(t *testing.T) {
		input := 1000000001
		
		result := isPalindrome(input)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	t.Run("handles negatives correctly", func(t *testing.T) {
		x := -121
		
		result := isPalindrome(x)
		
		assert.NotNil(t, result)
		assert.False(t, result)
		
	})
	t.Run("handles simple case correctly", func(t *testing.T) {
		x := 121
		
		result := isPalindrome(x)
		
		assert.NotNil(t, result)
		assert.True(t, result)
		
	})
}
