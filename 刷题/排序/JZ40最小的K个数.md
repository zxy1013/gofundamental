给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。 

```python
class Solution:
    def GetLeastNumbers_Solution(self, nums, k):
        if not nums:
            return []
        mid = nums[0]
        left = [x for x in nums[1:] if x < mid]
        right = [x for x in nums[1:] if x >= mid]
        if len(left)+1 == k:
            left.append(mid)
            return left
        # 还需要大于的排序
        elif len(left)+1 < k:
            left.append(mid)
            return left + self.GetLeastNumbers_Solution(right, k-len(left))
        # 小于的太多了
        else:
            return self.GetLeastNumbers_Solution(left, k)
```

```go
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
    if len(input) == 0{
        return []int{}
    }
    mid := input[0]
    left := make([]int,0)
    right := make([]int,0)
    for _,v := range input[1:]{
        if v > mid{
            right = append(right,v)
        }else{
            left = append(left,v)
        }
        
    }
    if len(left) + 1 == k{
        return append(left,mid)
    }else if len(left) + 1 < k{
        resu := make([]int,0)
        resu = append(left,mid)
        resu = append(resu,GetLeastNumbers_Solution(right,k-len(resu))...)
        return resu
    }else{
        return GetLeastNumbers_Solution(left,k)
    }
}
```

