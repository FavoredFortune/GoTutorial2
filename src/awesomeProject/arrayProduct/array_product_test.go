package arrayProduct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayProduct(t *testing.T){
	t.Run("it works for empty array", func(t *testing.T) {
		var empty []int
		
		result := arrayProduct(empty)
		
		assert.Nil(t, result)
	})
	
	t.Run("it works for example array", func(t *testing.T) {
		input := []int{2,3,4,5}
		
		result := arrayProduct(input)
		
		assert.NotNil(t, result)
		assert.Equal(t, []int{60,40,30,24}, result)
	})
	
	t.Run("it works if array has 0s", func(t *testing.T) {
		dangerZeros := []int{0,3,1}
		
		result := arrayProduct(dangerZeros)
		
		assert.NotNil(t, result)
		assert.Equal(t, []int{3, 0, 0}, result)
	})
	
	t.Run("it works with a single negative number in the array", func(t *testing.T) {
		negative := []int{-2, 3, 4, 5}
		
		result := arrayProduct(negative)
		fmt.Println(result)
		
		assert.NotNil(t, result)
		assert.Equal(t, []int{60, -40, -30, -24}, result)
		
	})
	
	t.Run("it works with many negative numbers in the array", func(t *testing.T) {
		negatives := []int{-2, 3, -4, -5}
		
		result := arrayProduct(negatives)
		fmt.Println(result)
		
		assert.NotNil(t, result)
		assert.Equal(t, []int{60, -40, 30, -24}, result)
		
	})
}