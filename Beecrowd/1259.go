package main
 
import (
    "fmt"
)
 
type Node struct {
	Value int 
	Right *Node 
	Left  *Node
}

func EvenAdd(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	} 
	if value < root.Value {
		root.Left = EvenAdd(root.Left, value)
	} else { 
		root.Right = EvenAdd(root.Right, value)
	}
	return root
}

func OddAdd(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	} 
	if value < root.Value {
		root.Right = OddAdd(root.Right, value)
	} else {
		root.Left = OddAdd(root.Left, value)
	}
	return root
}

func NeutralAdd(head *Node, value int) {
	if value % 2 == 0 {
		head.Left = EvenAdd(head.Left, value)
		return
	}
	head.Right = OddAdd(head.Right, value)
}

func PrintTree(root *Node) {
	if root == nil{
		return
	} 
	PrintTree(root.Left)
	fmt.Println(root.Value)
	PrintTree(root.Right)
}

func main() {
    n, x := 0, 0
		root := &Node{}
    fmt.Scanln(&n)
		for i := 0; i < n; i++ {
			fmt.Scanln(&x)
			NeutralAdd(root, x)
		}
		PrintTree(root.Left)
	PrintTree(root.Right)
}
