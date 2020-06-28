package recursion

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。


示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 调换
func swapPairs(head *ListNode) *ListNode {
	if nil == head {
		return head
	}

	var second = head.Next
	if nil == second {
		return head
	}

	// 转换动作，head当做2，second当做1

	third := second.Next

	fourth := swapPairs(third)

	head.Next = fourth
	second.Next = head

	return second
}

//
//func main() {
//	fourth := &ListNode{Val: 4, Next: nil}
//	third := &ListNode{Val: 3, Next: fourth}
//	second := &ListNode{Val: 2, Next: third}
//	head := &ListNode{Val: 1, Next: second}
//
//	newHead := swapPairs(head)
//	print(newHead)
//}
