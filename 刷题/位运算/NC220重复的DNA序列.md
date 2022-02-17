所有的 DNA 序列都是由 'A' , ‘C’ , 'G' , 'T' 字符串组成的，例如 'ACTGGGC' 。

编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。 

```go
package main
import "sort"

func repeatedDNA( DNA string ) []string {
    result := make(map[int]string,len(DNA))
    // 键：string 值：位置
    set := make(map[string]int,len(DNA))
    // 取键排序s
    key := make([]int,0)
    for i:=0;i<len(DNA)-9;i++{
        _,ok := set[DNA[i:i+10]]
        // 重复
        if ok {
            // 结果是否已经添加
            flag := true
            for _,v := range result{
                if v == DNA[i:i+10]{
                    flag = false
                }
            }
            if flag{
                result[i] = DNA[i:i+10]
                key = append(key,i)
            }
        }else{
            set[DNA[i:i+10]] = i
        }
    }
    sort.Ints(key)
    // 结果
    re := make([]string,0)
    for _,v := range key{
        re = append(re,result[v])
    }
    return re
}
```

```python
class Solution:
    def repeatedDNA(self , DNA: str) -> List[str]:
        flag = set()
        res = []
        for i in range(len(DNA)-9):
            if DNA[i:i+10] in flag and DNA[i:i+10] not in res:
                res.append(DNA[i:i+10])
            else:
                flag.add(DNA[i:i+10])
        return res
```

