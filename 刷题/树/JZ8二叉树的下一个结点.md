给定一个二叉树其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的next指针。 

```python
class Solution:
    def GetNext(self, pNode):
        # 该节点无右子树 且该节点为左子树
        if pNode.right == None and pNode.next and pNode.next.left == pNode:
            return pNode.next
        # 该节点无右子树 且该节点为右子树
        if pNode.right == None and pNode.next and pNode.next.right == pNode:
            x = pNode
            while x.next:
                if x == x.next.left:
                    break
                else:
                    x = x.next
            return x.next
        # 该节点有右子树
        if pNode.right != None:
            x = pNode.right
            while x.left:
                x = x.left
            return x
```

```go
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
    // 该节点有右子树
    if pNode.Right != nil{
        temp := pNode.Right
        for temp.Left != nil{
            temp = temp.Left
        }
        return temp
    }else{
        // 该节点无右子树
        // 该节点为左子树
        if pNode.Next != nil && pNode.Next.Left == pNode{
            return pNode.Next
        }
        // 该节点为右子树
        if pNode.Next != nil{
            temp := pNode.Next
            for temp != nil && temp.Next != nil && temp.Next.Left != temp{
                temp = temp.Next
            }
            return temp.Next
        }
    }
    return nil
}
```

