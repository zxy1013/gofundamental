在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
给定 target = 7，返回 true。
给定 target = 3，返回 false。 

```python
class Solution:
    def Find(self, target, array):
        # 二分 把二分值定在右上角
        j = len(array[0])-1
        i = 0 
        while i < len(array) and j >= 0:
            if array[i][j] == target:
                return True
            # 如果 tar > val, 说明target在更大的位置,说明第 i 行都是无效的，所以val下移
            if array[i][j] < target:
                i += 1
                continue
            # 如果 tar < val, 说明target在更小的位置，val右边的元素显然都是 > tar，所以val左移
            if array[i][j] > target:
                j -= 1
                continue
        return False
```

```go
func Find( target int ,  array [][]int ) bool {
    for i,j := 0,len(array[0])-1;i<len(array) && j>-1;{
        if array[i][j] == target{
            return true
        }else if array[i][j] > target{
            j -= 1
            continue
        }else{
            i += 1
            continue
        }
    }
    return false
}
```

 **NC86 矩阵元素查找**             

已知int一个有序矩阵**array**，同时给定矩阵的大小**n**和**m**以及需要查找的元素**target**，且矩阵的行和列都是从小到大有序的。设计查找算法返回所查找元素的二元数组，代表该元素的行号和列号(均从零开始)。保证元素互异。 

```go
func findElement( array [][]int ,  n int ,  m int ,  target int ) []int {
    for i,j := 0,m-1;i<n && j>-1;{
        if array[i][j] == target{
            return []int{i,j}
        }else if array[i][j] > target{
            j -= 1
            continue
        }else{
            i += 1
            continue
        }
    }
    return []int{-1,-1}
}
```