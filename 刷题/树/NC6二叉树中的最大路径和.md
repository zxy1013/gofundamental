二叉树里面的路径被定义为:从该树的任意节点出发，经过父=>子或者子=>父的连接，达到任意节点的序列。 

注意: 

1.同一个节点在一条二叉树路径里中最多出现一次 

2.一条路径至少包含一个节点，且不一定经过根节点 

> '''节点可能是非负的，因此开始dfs的节点不一定是根节点，结束的节点也不一定是叶子结点
>  题目没有说一定要按照自顶向下的顺序遍历，也就是说还有一种情况root.left->root->root.right。这就需要我们找到左子树最大值，右子树最大值加上根。'''

![F3A5B372D3918AF952A5EADE46E7CACD](F:\markdown笔记\刷题\树\F3A5B372D3918AF952A5EADE46E7CACD.png)

```python
class Solution:
    def maxPathSum(self , root ):
        import sys
        res = -sys.maxsize - 1 # python最小值
        def dfs(proot):
            nonlocal res
            if not proot:
                return 0 
            lmax = max(0,dfs(proot.left)) # 左子树的最大值 若为负 则为0
            rmax = max(0,dfs(proot.right)) # 右子树的最大值
            res = max(res , proot.val + lmax + rmax)
            return proot.val + max(lmax , rmax) # 找到root节点下的最大节点和 不找兄弟节点
        dfs(root)
        return res
```

```go
func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}
var maxu int
func dfs(root *TreeNode) int {
    if root == nil{
        return 0
    }
    lmax := max(0,dfs(root.Left))
    rmax := max(0,dfs(root.Right))
    maxu = max(maxu,lmax+rmax+root.Val)
    return root.Val + max(lmax,rmax)
}
func maxPathSum( root *TreeNode ) int {
    maxu = -int(0xffff)
    dfs(root)
    return maxu
}
```

