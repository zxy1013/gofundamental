现在有一个只包含数字的字符串，将该字符串转化成IP地址的形式，返回所有可能的情况。  

例如：  

给出的字符串为"25525522135",  

返回["255.255.22.135", "255.255.221.35"]. (顺序没有关系)  

```python
class Solution:
    def restoreIpAddresses(self , s ):
        res = []
        def dfs(temp,stri):
            if len(stri) == 0 and len(temp) == 4:
                res.append('.'.join(temp))
            if len(temp) < 4:
                for i in range(min(3,len(stri))):
                    p = stri[:i+1]
                    l = stri[i+1:]
                    # 每个地址包含4个十进制数，其范围为 0 - 255， 
                    # IPv4 地址内的数不会以 0 开头。
                    if p and 0 <= int(p) <= 255 and str(int(p)) == p:
                        temp.append(p)
                        dfs(temp,l)
                        temp.pop()
        dfs([],s)
        return res
```

```go
package main
import "strings"
import "strconv"

var res []string

func dfs(temp[]string,s string){
    if len(temp) == 4 && len(s) == 0{
        res = append(res, strings.Join(temp,"."))
    }
    if len(temp) < 4{
        for i:=0;i<3&&i<len(s);i++{
            left := s[:i+1]
            right := s[i+1:]
            strTint,_ := strconv.Atoi(left)
            // 不以0开头 但是有可能是0
            if strTint<256 && strTint>0 && left[0] != '0'|| len(left)==1 && left[0]=='0'{
                temp = append(temp,left)
                dfs(temp,right)
                temp = temp[:len(temp)-1]
            }
        }
    }
}
func restoreIpAddresses( s string ) []string {
    res = []string{}
    dfs([]string{},s)
    return res
}
```

