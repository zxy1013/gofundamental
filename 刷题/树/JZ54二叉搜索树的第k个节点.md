给定一棵结点数为n 二叉搜索树，请找出其中的第 k 小的TreeNode结点值。 

1.返回第k小的节点值即可 

2.不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1 

3.保证n个节点的值不一样 

```go
var ki int

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func dfs(proot *TreeNode) int {
    if proot == nil{
        return -1
    }
    l1 := dfs(proot.Left)
    ki -- 
    if ki == 0{
        return proot.Val
    }
    l2 := dfs(proot.Right)
    return max(l1,l2)
}

// 中序遍历
func KthNode( proot *TreeNode ,  k int ) int {
    ki = k
    return dfs(proot)
}
```

