# 关于单向链表
1、数据的前进是 ++2，链表的前进是list.Next
2、
```go
package ArrayList

type ListNode struct {
	Val  int
	Next *ListNode
}

var list,head *ListNode
list.Next.Next=head
head=head.Next
```
   则是将head赋值给别人，然后前进一步 
    

