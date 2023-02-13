package ArrayList

import (
	"math/rand"
)

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

type NumArray struct {
	sum []int
}

func Constructor1(nums []int) NumArray {
	l := len(nums)
	sum := make([]int, len(nums))
	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return NumArray{
		sum: sum,
	}
}
func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

//给你一个 下标从 0 开始 的正整数数组 w ，其中 w[i] 代表第 i 个下标的权重。
//请你实现一个函数 pickIndex ，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i 的 概率 为 w[i] / sum(w) 。
//例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。

type Solution struct {
	Sum    []int
	lenSum int
}

func Constructor2(w []int) Solution {
	var preSum []int
	for i := 1; i < len(w); i++ {
		preSum[i] = preSum[i-1] + w[i-1]
	}
	return Solution{
		preSum,
		len(preSum),
	}

}

func (s *Solution) PickIndex() int {
	l := s.Sum[len(s.Sum)-1]
	r := rand.Intn(l) + 1
	return LeftBound(s.Sum, r)
}

func LeftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r - l) / 2
		if target == nums[mid] {
			return mid
		} else if target >= nums[mid] {
			r--
		} else if target <= nums[mid] {
			l++
		}
	}
	return l + 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
