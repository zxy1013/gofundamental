以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

```python
class Solution:
    def solve(self , s , t ):
        s = s[::-1]
        t = t[::-1]
        carry = 0
        i = 0
        str1 = ''
        while i<len(s) or i<len(t) or carry:
            res = 0
            if i< len(s):
                res += int(s[i])
            if i<len(t):
                res += int(t[i])
            if carry:
                res += carry
            str1 = str(res % 10) + str1
            carry = res // 10
            i += 1
        return str1
```

```go
func solve( s string ,  t string ) string {
    carry := 0
    i := 0
    res := ""
    for i < len(s) || i < len(t)||carry != 0{
        temp := 0
        if i < len(s){
            temp += int(s[len(s)-1-i]-'0')
        }
        if i < len(t){
            temp += int(t[len(t)-1-i]-'0')
        }
        if carry != 0{
            temp += carry
        }
        carry = temp / 10
        res = strconv.Itoa(temp%10) + res
        i += 1
    }
    return res
}
```

```go
func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}
// 0-9 48-57
func solve( s string ,  t string ) string {
    carry := 0
    i := 0
    res := make([]byte,max(len(s),len(t))+1)
    for i < len(s) || i < len(t){
        temp := 0
        if i < len(s){
            temp += int(s[len(s)-1-i])-48
        }
        if i < len(t){
            temp += int(t[len(t)-1-i])-48
        }
        temp += carry
        carry = temp / 10
        res[len(res)-1-i] = byte((temp % 10) + 48)
        i += 1
    }
    if carry != 0{
        res[0] = '1'
        return string(res)
    }
    return string(res[1:])
}
```

```go
str := "12345"
temp := []byte(str)
// res := 18
// temp[2] -= res // invalid operation: temp[2] -= res (mismatched types byte and int)
temp[2] -= 18 // [49 50 33 52 53]
res1 := str[2] - 20 // 31
```