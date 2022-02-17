将给定的单链表  L： L0→L1→…→Ln−1→Ln 重新排序为：L0→Ln→L1→Ln−1→L2→Ln−2→…
1->2->3->4->5->6
第一步：将链表分为两个链表(快慢指针)
1->2->3  4->5->6
第二步：将第二个链表逆序
1->2->3  6->5->4
第二步：依次连接两个链表 连接形式如下
1->6->2->5->3->4

```python
class Solution:
    def reorderList(self , head ):
        if not head or not head.next:
            return head
        def reverse(head1):
            temp = None
            while head1:
                p = head1
                head1 = head1.next
                p.next = temp
                temp = p
            return temp
        # 第一步 求中间节点slow
        slow , fast = head , head
        while fast and fast.next and fast.next.next:
            slow = slow.next
            fast = fast.next.next
        # 第二步 逆序slow
        newhead = reverse(slow.next)
        slow.next = None
        result = head
        # 第三步 连接两个链表
        while newhead:
            temp = head.next
            head.next = ListNode(newhead.val)
            newhead = newhead.next
            head.next.next = temp
            head = head.next.next
        return result
```

```go
func reverse(head *ListNode)*ListNode{
    var newhead *ListNode
    for head != nil{
        temp := head
        head = head.Next
        temp.Next = newhead
        newhead = temp
    }
    return newhead
}
func reorderList( head *ListNode ) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    slow,fast := head,head
    for fast != nil && fast.Next != nil && fast.Next.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    result := head
    new := reverse(slow.Next)
    slow.Next = nil
    for new != nil && head != nil{
        temp := head.Next
        head.Next = &ListNode{Val:new.Val}
        new = new.Next
        head.Next.Next = temp
        head = head.Next.Next
    }
    return result
}
```

