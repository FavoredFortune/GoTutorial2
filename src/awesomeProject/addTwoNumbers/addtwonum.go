package addTwoNumbers

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	
	result := &ListNode{}
	current := result
	carry := 0
	
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		
		if l1 !=nil{
			sum += l1.Val
			l1 = l1.Next
		}
		
		if l2 !=nil{
			sum += l2.Val
			l2 = l2.Next
		}
	
	}
	
	return &result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	
	for l1.Next !=nil || l2.Next !=nil {
		sum := l1.Val + l2.Val
		a := sum / 10
		result.Next = &ListNode{Val: a}
		result = result.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	
	return result
}