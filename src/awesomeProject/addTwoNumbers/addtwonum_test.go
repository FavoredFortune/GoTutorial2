package addTwoNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T){
	t.Run("addTwoNumbers works", func(t *testing.T) {
		l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
		l2 := &ListNode {5, &ListNode{6, &ListNode{4, nil}}}
		expected := &ListNode{7, &ListNode{0, &ListNode{8,nil}}}
		
		result := addTwoNumbers(l1, l2)
		assert.NotNil(t, result)
		assert.Equal(t,expected, result )
	})
	
	t.Run("addTwoNumbers works with negative numbers", func(t *testing.T) {
		l1 := &ListNode{-2, &ListNode{-4, &ListNode{3, nil}}}
		l2 := &ListNode {5, &ListNode{6, &ListNode{4, nil}}}
		expected := &ListNode{3, &ListNode{2, &ListNode{7,nil}}}
		
		result := addTwoNumbers(l1, l2)
		assert.NotNil(t, result)
		assert.Equal(t,expected, result )
	})
	
	
	t.Run("addTwoNumbers works with 0s", func(t *testing.T) {
		l1 := &ListNode{0, &ListNode{0, &ListNode{0, nil}}}
		l2 := &ListNode {0, &ListNode{0, &ListNode{0, nil}}}
		expected := &ListNode{0, &ListNode{0, &ListNode{0,nil}}}
		
		result := addTwoNumbers(l1, l2)
		assert.NotNil(t, result)
		assert.Equal(t,expected, result )
	})
}
