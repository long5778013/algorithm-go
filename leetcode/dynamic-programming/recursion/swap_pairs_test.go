package recursion

import "testing"

func Test_SwapPairs(t *testing.T) {
	fourth := &ListNode{Val: 4, Next: nil}
	third := &ListNode{Val: 3, Next: fourth}
	second := &ListNode{Val: 2, Next: third}
	head := &ListNode{Val: 1, Next: second}

	newHead := swapPairs(head)
	print(newHead)
	if newHead.Val == 2 && newHead.Next.Val == 1 && newHead.Next.Next.Val == 4 && newHead.Next.Next.Next.Val == 3 {
		t.Log("swapPairs success")
	} else {
		t.Error("swapPairs failed")
	}
}
