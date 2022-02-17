 输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。 

```python
class Solution:
    def printListFromTailToHead(self , listNode: ListNode) -> List[int]:
        # 递归
        res = []
        if not listNode:
            return []
        def dfs(root):
            if not root.next:
                res.append(root.val)
                return 
            else:
                dfs(root.next)
            res.append(root.val)
        dfs(listNode)
        return res
    
        # 反向输出
        res=[]
        while listNode:
            res.append(listNode.val)
            listNode = listNode.next
        return res[::-1]
    
        # 反转链表
        Head = listNode
        temp = None
        res = []
        while Head:
            # x记录head节点
            x = Head
            Head = Head.next
            x.next = temp
            temp = x
        while temp:
            res.append(temp.val)
            temp = temp.next
        return res
```

```go
var res []int

func dfs(head *ListNode){
    if head.Next == nil{
        res = append(res,head.Val)
    }else{
        dfs(head.Next)
        res = append(res,head.Val)
    }
}

func printListFromTailToHead( head *ListNode ) []int {
    // 递归
    if head == nil{
        return []int{}
    }
    dfs(head)
    return res
}


// 递归的匿名函数版
func printListFromTailToHead( head *ListNode ) []int {
    ans := []int{}
    var f func(head *ListNode )
    f = func(head *ListNode ){
        if head == nil {
           return
        }
        f(head.Next)
        ans = append(ans, head.Val)
    }
    f(head)
    return ans
}


// 反转链表
func printListFromTailToHead( head *ListNode ) []int {
   var tail *ListNode
    for head != nil{
        temp := head
        head = head.Next
        temp.Next = tail
        tail = temp
    }
    var res []int
    for tail != nil{
        res = append(res,tail.Val)
        tail = tail.Next
    }
    return res
}
```

