package sortList

// Definition for singly-linked list.
 type ListNode struct {
    Val int
     Next *ListNode
 }

func sortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var tempLow *ListNode
	var tempHigh *ListNode
	var temp *ListNode
	
	for head.Next != nil{
		if head.Val < head.Next.Val{
			temp = head.Next
			head.Next = head
			head = temp
		}
		temp = head
		temp.Next = head.Next.Next
		head = head.Next
		head.Next = temp
		head = head.Next
	}
	return head
}

// sort problem!!! go review merge sort! see examples:
//https://leetcode.com/explore/interview/card/top-interview-questions-hard/117/linked-list/840/discuss/136335/merge-sort-using-go
//https://leetcode.com/problems/sort-list/discuss/298786/Go-Merge-sort.-O(1)-space
//Extra memory solution
//https://leetcode.com/problems/sort-list/discuss/287418/go-simple-code-need-extra-memory

//Java QuickSort MergeSort
//https://leetcode.com/problems/sort-list/discuss/361447/Java-quick-sort-and-merge-sort

//Java
//https://leetcode.com/problems/sort-list/discuss/46714/Java-merge-sort-solution