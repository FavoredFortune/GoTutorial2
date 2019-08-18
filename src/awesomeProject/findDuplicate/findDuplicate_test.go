package findDuplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T){
	t.Run("it finds leetcode duplicate test 1", func(t *testing.T) {
		input := []int{1,3,4,2,2}
		
		result := findDuplicate(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 2, result)
	})
	
	t.Run("it finds leetcode duplicate test 2", func(t *testing.T) {
		input := []int{3,1,3,4,2}
		
		result := findDuplicate(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, 3, result)
	})
}
