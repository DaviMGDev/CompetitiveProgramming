/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l3 := &ListNode{
        Val: 0,
        Next: nil,
    }

    a, b, c := l1, l2, l3
    carry := 0
    sum   := 0
    for a != nil || b != nil || carry != 0 {

        x := 0
        if a != nil {
            x = a.Val
            a = a.Next
        }
        y := 0
        if b != nil {
            y = b.Val
            b = b.Next
        }
        
        sum = x + y + carry
        if sum != 0 {
            c.Next = &ListNode{}
            c = c.Next
            value := sum % 10
            c.Val = value
        }
        carry = sum / 10
    }

    return l3
}
