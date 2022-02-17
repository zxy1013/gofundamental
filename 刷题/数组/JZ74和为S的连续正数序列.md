小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,他马上就写出了正确答案是100。但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。现在把问题交给你,你能不能也很快的找出所有和为S的连续正数序列?  

> 初始化，i=1,j=1, 表示窗口大小为0
> 如果窗口中值的和小于目标值sum， 表示需要扩大窗口，j += 1
> 否则，如果大于目标值sum，表示需要缩小窗口，i += 1 

```go
package main

func FindContinuousSequence( sum int ) [][]int {
    res := make([][]int,0)
    i := 1 // 左边界
    j := 1 // 右边界
    temp := 0 // 当前窗口的sum
    // 窗口左边界走到sum的一半即可终止，因为题目要求至少包含2个数 
    for i<= sum/2{
        if temp == sum{
            templist := make([]int,0)
            // 不包括右边界
            for k:=i;k<j;k++{
                templist = append(templist,k)
            }
            res = append(res, templist)
            temp -= i
            i++
        }else if temp > sum{
            temp -= i
            i ++
        }else{
            temp += j
            j++
        }
    }
    return res
}
```

