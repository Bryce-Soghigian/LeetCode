/**
func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
 */
func middleNode(head *ListNode) *ListNode {
    // a, b
    // 1, 2, 3, 4 , 5
    //       a
    //             b
    slow, fast := head, head

    for fast != nil && fast.Next != nil{
        slow = slow.Next 
        fast = fast.Next.Next
    }
    return slow
}