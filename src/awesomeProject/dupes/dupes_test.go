package dupes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainesDupes(t *testing.T){
	t.Run("it finds duplicates", func(t *testing.T) {
		input := []int{1,2,2,3,5,6}
		
		result := containsDuplicate(input)
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	
	t.Run("it returns false when there are no duplicates", func(t *testing.T) {
		input := []int{6,4,2,3}
		
		result := containsDuplicate(input)
		assert.NotNil(t, result)
		assert.False(t, result)
	})
	
	t.Run("it returns false when array input is too short to have duplicates", func(t *testing.T) {
		input := []int{4}
		
		result := containsDuplicate(input)
		assert.NotNil(t, result)
		assert.False(t, result)
	})
}
