package ArrayList

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转整个列表，使用递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 反转链表前N个
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		D = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = D
	return last
}

// 提供一个驱动节点，接收last
var D *ListNode

// 反转m-n区间
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}
