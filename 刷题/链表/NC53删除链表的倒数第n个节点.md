删除链表的倒数第n个节点 
给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
给出的链表为: 1→2→3→4→5, n=2，删除了链表的倒数第 n 个节点之后,链表变为1→2→3→5.

> 可以使用两个指针fast 和 slow 同时对链表进行遍历，并且 fast 比 slow 超前 n 个节点。
>         当 fast 遍历到链表的末尾时，slow 就恰好处于倒数第 n 个节点。
>         初始时 fast 和 slow 均指向头节点。首先使用 fast 对链表进行遍历，遍历的次数为 n。此时，fast 和 slow 之间间隔了 n-1 个节点，
>         即 fast 比 slow 超前了 n 个节点。在这之后，同时使用 fast 和 slow 对链表进行遍历。
>         当 fast 遍历到链表的末尾（即 fast 为空指针）时，slow 恰好指向倒数第 n 个节点。
>

```python
class Solution:
    def removeNthFromEnd(self , head , n ):
        slow,fast = head,head
        while n :
            fast = fast.next
            n -= 1
        if not fast: # 判断n是否是头结点
            return head.next
        while fast.next:
            slow = slow.next
            fast = fast.next
        slow.next = slow.next.next
        return head
```

```go
func removeNthFromEnd( head *ListNode ,  n int ) *ListNode {
    fast,slow := head,head
    for n > 0 && fast != nil {
        n -- 
        fast = fast.Next
    }
    if fast == nil{
        return head.Next
    }else{
        for fast.Next != nil{
            slow = slow.Next
            fast = fast.Next
        }
        slow.Next = slow.Next.Next
        return head
    }
}
```

