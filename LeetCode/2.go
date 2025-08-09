/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b  := l1, l2
	head  := &ListNode{}
	c     := head
	carry := 0
	x, y  := 0, 0
	//for a != nil || b != nil || carry != 0 {
    for c != nil {
		x = 0
		if a != nil {
			x = a.Val
			a = a.Next
		}
		y = 0
		if b != nil {
			y = b.Val
			b = b.Next 
		}
		c.Val = x + y + carry
		carry = c.Val / 10 
		c.Val = c.Val % 10
        if a!= nil || b != nil || carry != 0 {
            c.Next = &ListNode{}
        }
        c = c.Next
	}
	return head
}
