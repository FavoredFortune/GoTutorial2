package removeDupesSorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDupesSorted(t *testing.T){
	t.Run("returns length of deduped sorted array", func(t *testing.T) {
		test2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		
		result := removeDuplicates(test2)
		
		assert.NotNil(t, result)
		assert.Equal(t, 5, result)
	})
	
	t.Run("returns length of array when nothing duplicated", func(t *testing.T) {
		testArr := []int{1, 2}
		
		result := removeDuplicates(testArr)
		
		assert.NotNil(t, result)
		assert.Equal(t, 2, result)
	})
}