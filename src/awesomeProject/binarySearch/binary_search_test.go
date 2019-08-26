package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiSearch(t *testing.T) {
	t.Run("it finds the indicies of the target", func(t *testing.T) {
		nums := []int {1, 2, 3, 4, 5, 7}
		target := 7
		
		result := BiSearch(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 5, result)
	})
	
	t.Run("it finds the indicies of a negative value target", func(t *testing.T) {
		nums := []int {-1, 0, 1, 2, 3, 4, 5, 7, 33}
		target := -1
		
		result := BiSearch(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 0, result)
	})

	t.Run("it returns -1 if the target isn't in the sorted slice", func(t *testing.T) {
		nums := []int {1, 2, 3, 4, 5, 7}
		target := 6
		
		result := BiSearch(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, -1  , result)
	})
}

func TestFindPosition(t *testing.T) {
	t.Run("it returns the position where the value belongs in the sorted slice", func(t *testing.T) {
		nums := []int {-10, -4, 0, 3, 7, 9, 12, 42}
		target := 5
		
		result := FindPosition(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 4  , result)
	})
	
	t.Run("it returns the position where the value belongs in the sorted slice when it belongs at the end",
		func(t *testing.T) {
		nums := []int {-10, -4, 0, 3, 7, 9, 12, 42}
		target := 55
		
		result := FindPosition(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 8  , result)
	})
	
	t.Run("it returns the position where the value belongs in the sorted slice when slice is only 1 int long",
		func(t *testing.T) {
			nums := []int { 3}
			target := 5
			
			result := FindPosition(nums, target)
			
			assert.NotNil(t, result)
			assert.Equal(t, 1  , result)
		})
	
	t.Run("it returns the position where the value belongs in the sorted slice when target value is negative",
		func(t *testing.T) {
			nums := []int { -10, -4, 0, 3, 7, 9, 12, 42}
			target := -15
			
			result := FindPosition(nums, target)
			
			assert.NotNil(t, result)
			assert.Equal(t, 0  , result)
		})
	
	t.Run("it returns the position where the value belongs in the sorted slice when the slice is empty",
		func(t *testing.T) {
			 var nums []int
			target := -15
			
			result := FindPosition(nums, target)
			
			assert.NotNil(t, result)
			assert.Equal(t, 0  , result)
		})
}

func TestVerifyBinary(t *testing.T) {
	t.Run("it returns true when slice is binary sorted", func(t *testing.T) {
		nums := []int {0, 1,2, 3}
		
		result := VerifyBinary(nums)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
	
	t.Run("it returns false when slice is not binary sorted", func(t *testing.T) {
		nums := []int {0, 12, 3, 7}
		
		result := VerifyBinary(nums)
		
		assert.NotNil(t, result)
		assert.False(t, result)
	})
	
	t.Run("it returns true when slice is empty (effectively sorted)", func(t *testing.T) {
		var nums []int
		
		result := VerifyBinary(nums)
		
		assert.NotNil(t, result)
		assert.True(t, result)
	})
}

func TestReturnNewBinarySort(t *testing.T) {
	t.Run("correctly inserts target value into slice that is returned sorted", func(t *testing.T) {
		nums := []int {0,1,2,4,5}
		target := 3
		
		result := ReturnNewBinarySort(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 6 ,len(result))
		assert.ElementsMatch(t, []int{0,1,2,3,4,5}, result)
	})
	
	t.Run("correctly inserts negative target value into slice that is returned sorted", func(t *testing.T) {
		nums := []int {-1, 0,1,2,4,5}
		target := -3
		
		result := ReturnNewBinarySort(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 7 ,len(result))
		assert.ElementsMatch(t, []int{-3, -1,0,1,2,4,5}, result)
	})
	t.Run("it does not error when inserting a value into an empty slice", func(t *testing.T) {
		var nums []int
		target := 42
		
		result := ReturnNewBinarySort(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 1 ,len(result))
		assert.ElementsMatch(t, []int{42}, result)
		
	})
	t.Run("it does correctly inserts a value when the value already exists", func(t *testing.T) {
		nums := []int {0,1,2,3,4,5}
		target := 3
		
		result := ReturnNewBinarySort(nums, target)
		
		assert.NotNil(t, result)
		assert.Equal(t, 7 ,len(result))
		assert.ElementsMatch(t, []int{0,1,2,3,3,4,5}, result)
		
	})
}