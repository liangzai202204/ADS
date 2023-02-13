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

3、关于快慢指针，用快慢指针寻找链表的中点，如果fast指针没有指向null，说明链表长度为奇数，slow还要再前进一步
```go
if (fast != null)
slow = slow.next;
```
4、

