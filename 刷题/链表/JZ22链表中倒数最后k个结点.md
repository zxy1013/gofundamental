输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数后k个节点。 如果该链表长度小于k，请返回一个长度为 0 的链表。 

**快慢指针** 第一个指针先移动k步，然后第二个指针再从头开始，这个时候这两个指针同时移动，当第一个指针到链表的末尾的时候，返回第二个指针即可

```python
class Solution:
    def FindKthToTail(self , pHead , k ):
        if not pHead:
            return None
        # 双指针
        fast,slow = pHead,pHead
        # 先移动k步
        for i in range(k):
            if fast:
                fast = fast.next
            else:
                return None
        while fast:
            fast = fast.next
            slow = slow.next
        return slow
```

```go
func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
    fast,slow := pHead,pHead
    // 先移动k步
    for fast != nil && k > 0{
        k--
        fast = fast.Next
    }
    if k > 0{
        return nil
    }
    for fast != nil{
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
```

