默认情况下，Go 语言使用的是值传递，指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。使用普通变量作为函数参数的时候，在传递参数时只是对变量值拷贝，即将实参的值复制给变参，当函数对变参进行处理时，并不会影响原来实参的值。

```go
func test1(a []int){
	a[1] = 1
}

func test2(a []int){
	a = append(a,5)
	fmt.Println(a,len(a),cap(a)) // [0 1 4 5] 4 4
}

func main(){
	arr := make([]int,0)
	arr = append(arr,0)
	arr = append(arr,2)
	arr = append(arr,4)
	fmt.Println(arr,len(arr),cap(arr)) // [0 2 4] 3 4
	test1(arr)
	fmt.Println(arr,len(arr),cap(arr)) // [0 1 4] 3 4
	test2(arr)
	fmt.Println(arr,len(arr),cap(arr)) // [0 1 4] 3 4

	// 相当于函数内部进行了值拷贝，共享一个底层数组，若是函数内部扩容了，则不共享。
	arr1 := make([]int,0)
	arr1 = append(arr1,0)
	arr1 = append(arr1,2)
	arr1 = append(arr1,4)
	fmt.Println(arr1,len(arr1),cap(arr1)) // [0 2 4] 3 4
	arr2 := arr1
	arr2[1] = 1
	fmt.Println(arr1,len(arr1),cap(arr1)) // [0 1 4] 3 4
	arr2 = append(arr2, 5)
	fmt.Println(arr2,len(arr2),cap(arr2)) // [0 1 4 5] 4 4
	fmt.Println(arr1,len(arr1),cap(arr1)) // [0 1 4] 3 4
}
```



python不允许程序员选择采用传值还是传引用。Python参数传递采用的肯定是“传对象引用”的方式。这种方式相当于传值和传引用的一种综合。如果函数收到的是一个可变对象（比如字典或者列表）的引用，就能修改对象的原始值-----相当于通过“传引用”来传递对象。如果函数收到的是一个不可变对象（比如数字、字符或者元组）的引用，就不能直接修改原始对象-----相当于通过“传值'来传递对象。 



```python
class Solution:
    def reConstructBinaryTree(self, pre, vin):
        if not vin or not pre:
            return None
        proot = TreeNode(pre[0]) # 根节点
        indexr = vin.index(pre.pop(0))
        pleft = vin[:indexr] # 左子树
        pright = vin[indexr+1:] # 右子树
        # 引用传递
        proot.left = self.reConstructBinaryTree(pre,pleft)
        # 上下两个pre是一个引用
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
    // 值传递
    proot.Left = reConstructBinaryTree(pre[:index],vin[:index])
    // 上下两个pre是同一个值
    proot.Right = reConstructBinaryTree(pre[index:],vin[index+1:])
    """
    proot.Left = reConstructBinaryTree(pre[:],left)
    proot.Right = reConstructBinaryTree(pre[:],right)
    """
    return proot
}
```
