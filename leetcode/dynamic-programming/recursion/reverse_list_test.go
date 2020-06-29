package recursion

import (
	"fmt"
	"testing"
)

func TestReverseListIteration(t *testing.T) {
	head := new(ListNode)
	head.Val = 1
	n2 := new(ListNode)
	n2.Val = 2
	n3 := new(ListNode)
	n3.Val = 3
	n4 := new(ListNode)
	n4.Val = 4
	n5 := new(ListNode)
	n5.Val = 5
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	prev := reverseListIteration(head)
	fmt.Println(prev)

	if prev.Val == 5 && prev.Next.Val == 4 {
		t.Log("TestReverseList success")
	}
}
func TestReverseListRecursion(t *testing.T) {
	head := new(ListNode)
	head.Val = 1
	n2 := new(ListNode)
	n2.Val = 2
	n3 := new(ListNode)
	n3.Val = 3
	n4 := new(ListNode)
	n4.Val = 4
	n5 := new(ListNode)
	n5.Val = 5
	head.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	prev := reverseListRecursion(head)
	fmt.Println(prev)

	if prev.Val == 5 && prev.Next.Val == 4 {
		t.Log("TestReverseList success")
	}
}
