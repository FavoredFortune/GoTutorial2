package twoSum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T){
	nums := []int {2,7,11,15}
	target := 9
	
	t.Run("it finds the target and returns the correct indicies", func(t *testing.T) {
		result := TwoSum(nums, target)
		expected := []int{0,1}
		
		assert.NotNil(t, result)
		assert.Equal(t,expected, result)
		fmt.Println(expected, result)
		
	})
	
	t.Run("it correctly returns when input array is empty", func(t *testing.T) {
		result := TwoSum([]int{}, target)
		expected := []int{}
		
		assert.NotNil(t, result)
		assert.Equal(t,expected, result)
		fmt.Println(expected, result)
		
	})
	
	t.Run("works as expected when target isn't in domain", func(t *testing.T) {
		result := TwoSum(nums, 3)
		expected := []int{}
		
		assert.NotNil(t, result)
		assert.Equal(t,expected, result)
		fmt.Println(expected, result)
		
	})
	
	t.Run("works when target is negative", func(t *testing.T) {
		result := TwoSum([]int{-5, 2, 11, 12}, -3)
		expected := []int{0, 1}
		
		assert.NotNil(t, result)
		assert.Equal(t,expected, result)
		fmt.Println(expected, result)
	})
	
}

//fmt.Println(twoSum(nums, target))
//
//nums2 := []int {3,2,4}
//target2 := 6
//fmt.Println(twoSum(nums2, target2))
//
//nums3 := []int {3,2,3}
//target3 := 6
//fmt.Println(twoSum(nums3, target3))