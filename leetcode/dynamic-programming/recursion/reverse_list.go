package recursion

// 反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 调换动作传递
func reverseList(head *ListNode) *ListNode {
	// 1前为nil
	var prev *ListNode = nil
	var next *ListNode = nil
	// 顺次迭代，到nil停止
	for head != nil {
		// 移动迭代指针
		next = head.Next
		// 当前节点下一节点，反向指
		head.Next = prev
		// 下次迭代，当前节点变为前节点
		prev = head
		// 移动到下一节点
		head = next
	}
	return prev
}
