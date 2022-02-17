链表中的节点每k个一组翻转 
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表，如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样

```python
class Solution:
    def reverseKGroup(self , head , k ):
        def reverse(a,b): # 翻转
            tem = None
            while a != b:
                s = a
                a = a.next
                s.next = tem
                tem = s
            return tem # 头结点
        
    	b = head
        for i in range(k): # 不足k不翻转
             if not b:
                 return head
             b = b.next
         newhead = reverse(head,b)
         head.next = self.reverseKGroup(b , k ) # 此时head是尾
         return newhead
```

```go
func reverse(head *ListNode,tail *ListNode) *ListNode{
    var newhead *ListNode
    for head != tail{
        temp := head
        head = head.Next
        temp.Next = newhead
        newhead = temp
    }
    return newhead
}
func reverseKGroup( head *ListNode ,  k int ) *ListNode {
    tail := head
    for i:=0;i<k;i++{
        if tail == nil{
            return head
        }else{
            tail = tail.Next
        }
    }
    newrhead := reverse(head,tail)
    head.Next = reverseKGroup(tail,k)
    return newrhead
}
```

两两交换链表的节点             

一个链表，两两交换相邻节点，需要真正交换节点本身，而不是修改节点的值。 两两交换示例： 链表：1->2->3->4  交换后：2->1->4->3  链表：1->2->3  交换后：2->1->3

```python
class Solution:
    def swapLinkedPair(self , head: ListNode) -> ListNode:
        def reverse(a,b): # 翻转
            tem = None
            while a != b:
                s = a
                a = a.next
                s.next = tem
                tem = s
            return tem # 头结点
        b = head
        for i in range(2): # 不足k不翻转
             if not b:
                return head
             b = b.next
        newhead = reverse(head,b)
        head.next = self.swapLinkedPair(b) # 此时head是尾
        return newhead
```

```python
class Solution:
    def swapLinkedPair(self , head: ListNode) -> ListNode:
        if not head or not head.next:
            return head
        cur = head
        newhead = ListNode(0)
        result = newhead
        while cur and cur.next:
            temp = cur.next
            newhead.next = temp
            cur.next = cur.next.next
            temp.next = cur
            cur = cur.next
            newhead = newhead.next.next
        return result.next
```

```go
func swapLinkedPair( head *ListNode ) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    // 虚拟头结点
    newhead := &ListNode{Val: 0}
    result := newhead
    cur := head
    for cur != nil && cur.Next != nil{
        temp := cur.Next
        cur.Next = cur.Next.Next
        temp.Next = cur
        newhead.Next = temp
        newhead = temp.Next
        cur = cur.Next
    }
    return result.Next
}
```

