给定节点数为 n 二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。 

> """
>         根据前序序列第一个结点确定根结点
>         根据根结点在中序序列中的位置分割出左右两个子序列
>         对左子树和右子树分别递归使用同样的方法继续分解 """

```python
class Solution:
    def reConstructBinaryTree(self, pre, vin):
        if not vin or not pre:
            return None
        proot = TreeNode(pre[0]) # 根节点
        indexr = vin.index(pre.pop(0))
        pleft = vin[:indexr] # 左子树
        pright = vin[indexr+1:] # 右子树
        proot.left = self.reConstructBinaryTree(pre,pleft)
        proot.right = self.reConstructBinaryTree(pre,pright)
        return proot
```

```go
func findindex(a[]int,in int) int{
    for i, v :=  range a {
        if v == in{
            return i
        }
    }
    return -1
}

func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
    if len(pre) == 0 || len(vin) == 0{
        return nil
    }
    proot := &TreeNode{Val: pre[0]}
    index:= findindex(vin,pre[0])
    pre = pre[1:]
    proot.Left = reConstructBinaryTree(pre[:index],vin[:index])
    proot.Right = reConstructBinaryTree(pre[index:],vin[index+1:])
    return proot
}
```

给定一个二叉树的中序与后序遍历结果，请你根据两个序列构造符合这两个序列的二叉树。  

```python
class Solution:
    def buildTree(self , inorder: List[int], postorder: List[int]) -> TreeNode:
        if not inorder or not postorder:
            return None
        proot = TreeNode(postorder[-1]) # 根节点
        indexr = inorder.index(postorder.pop(-1))
        pleft = inorder[:indexr] # 左子树
        pright = inorder[indexr+1:] # 右子树
        proot.right = self.buildTree(pright,postorder)
        proot.left = self.buildTree(pleft,postorder)
        return proot
```

```go
func findindex(a[]int,in int) int{
    for i, v :=  range a {
        if v == in{
            return i
        }
    }
    return -1
}

func buildTree( inorder []int ,  postorder []int ) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0{
        return nil
    }
    proot := &TreeNode{Val: postorder[len(postorder)-1]}
    index:= findindex(inorder,postorder[len(postorder)-1])
    postorder = postorder[:len(postorder)-1]
    proot.Right = buildTree(inorder[index+1:],postorder[index:])
    proot.Left = buildTree(inorder[:index],postorder[:index])
    return proot
}
```

