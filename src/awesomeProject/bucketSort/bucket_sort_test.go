package bucketSort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucketSort(t *testing.T){
	t.Run("it sorts small amount of numbers without duplicates", func(t *testing.T) {
		nums := []int {1, 10, 5, 2, 20}
		
		result, err := Sort(nums)
		
		fmt.Print(result)
		
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, []int{1,2,5,10,20}, result)
	})
	
	t.Run("it handles an empty input", func(t *testing.T) {
		var nums[]int
		
		result, err := Sort(nums)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
	
	t.Run("it sorts numbers with duplicates", func(t *testing.T) {
		nums := []int {1, 10, 5, 2, 20, 2, 10}
		
		result, err := Sort(nums)
		
		fmt.Print(result)
		
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, []int{1,2,2,5,10,10,20}, result)
	})
}
