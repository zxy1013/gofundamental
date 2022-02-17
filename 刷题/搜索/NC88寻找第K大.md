有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。 

给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。 

```python
class Solution:
    def findKth(self, a, n, K):
        def quicksort(a,k):
            if not a:
                return None
            mid = a[0]
            left = [x for x in a[1:] if x>=a[0]]
            right = [x for x in a[1:] if x<a[0]]
            if len(left) == k-1:
                return mid
            elif len(left) > k-1:
                return quicksort(left, k)
            else:
                return quicksort(right, k-len(left)-1)
        return quicksort(a, K)
```

```go
package main

func dfs(a []int,target int)int{
    mid := a[0]
    left := make([]int,0)
    right := make([]int,0)
    // 第k大 left大于 right小于
    for i,_ := range a[1:]{
        if a[i+1] < mid{
            right = append(right,a[i+1])
        }else{
            left = append(left,a[i+1])
        }
    }
    if len(left) == target - 1{
        return mid
    }else if len(left) > target - 1{
        return dfs(left,target)
    }else if len(left) < target - 1{
        // -1 表示mid
        return dfs(right,target-len(left)-1)
    }
    return -1
}
func findKth( a []int ,  n int ,  K int ) int {
    return dfs(a,K)
}
```

