在一个长度为n的数组里的所有数字都在0到n-1的范围内。  数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任一一个重复的数字。  例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，那么对应的输出是2或者3。存在不合法的输入的话输出-1 

```python
class Solution:
    def duplicate(self , numbers ):
        for i in range(len(numbers)):
            if numbers[i] in numbers[:i]:
                return numbers[i]
        return -1
```

  

```go
func sort(numbers []int)[]int{
    if len(numbers) <= 1{
        return numbers
    }
    left := make([]int,0)
    right := make([]int,0)
    result:= make([]int,0)
    for _,v := range numbers[1:]{
        if v > numbers[0]{
            right = append(right,v)
        }else{
            left = append(left,v)
        }
        
    }
    result = append(result,sort(left)...)
    result = append(result,numbers[0])
    result = append(result,sort(right)...)
    return result
}
func duplicate( numbers []int ) int {
    num := sort(numbers)
    for i:=1;i<len(num);i++{
        if num[i] == num[i-1]{
            return num[i]
        }
    }
    return -1
}
```

