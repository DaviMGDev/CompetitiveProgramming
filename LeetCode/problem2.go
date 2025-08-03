package main 

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func (ll *ListNode) Append(value int) {
	newNode := &ListNode{Val: value, Next: nil}
	if ll == nil {
		return newNode
	}
	i := ll
	for i.Next != nil{
		i = i.Next
	}
	i.Next = newNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := l1, l2
	l3    := &ListNode{Val: 0, Next: nil}
	c     := l3
	carry, sum := false, 0
	for {
		if a == nil {
			if b == nil {
				if carry {
					l3.Append(1)
				}
				return l3
			} else {
				if carry {
					sum
				}
			}
		} else {
			if b == nil {
			} else {

			}
		}
	}
}

func main() {

}
