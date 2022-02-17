给定一个二叉树的根节点root，该树的节点值都在数字 0−9之间，每一条从根节点到叶子节点的路径都可以用一个数字表示。 

1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

```python
class Solution:
    def sumNumbers(self , root ):
        if not root:
            return 0
        res = 0
        def dfs(root,temp):
            nonlocal res
            if not root.left and not root.right:
                res += temp
            if root.left:
                dfs(root.left,temp*10+root.left.val)
            if root.right:
                dfs(root.right,temp*10+root.right.val)
        dfs(root,root.val)
        return res
```



```go
func sumNumbers( root *TreeNode ) int {
    var result int
    var f func(root *TreeNode,sum int)
    f = func(root *TreeNode,sum int){
        if root.Left == nil && root.Right == nil{
            result += sum
        }
        if root.Left != nil{
            f(root.Left,sum*10 + root.Left.Val)
        }
        if root.Right != nil{
            f(root.Right,sum*10 + root.Right.Val)
        }
    }
    if root != nil{
        f(root,root.Val)
    }
    return result
}
```

