package addTwoNumbers

type ListNode struct{
	Val int
	Next *ListNode
}

// help with how to move/create nodes in go with https://blog.csdn.net/a32521500/article/details/77650467

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
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
		
		carry = sum /10
		current.Next = new(ListNode)
		current.Next.Val = sum %10
		current = current.Next
	
	}
	
	return result.Next
}

