写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。传入的字符串可能有以下部分组成:

> 1.若干空格
>
> 2.（可选）一个符号字符（'+' 或 '-'）
>
> 3.数字，字母，符号，空格组成的字符串表达式
>
> 4.若干空格

> 转换算法如下:
> 1.去掉无用的前导空格
> 2.第一个非空字符为+或者-号时，作为该整数的正负号，如果没有符号，默认为正数
> 3.判断整数的有效部分：
>   3.1 确定符号位之后，与之后面尽可能多的连续数字组合起来成为有效整数数字，如果没有有效的整数部分，那么直接返回0
>   3.2 将字符串前面的整数部分取出，后面可能会存在多余的字符(字母，符号，空格等)，这些字符可以被忽略，它们对于函数不应该造成影响
>   3.3  整数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31的整数应该被调整为 −2^31 ，大于 2^31 − 1 的整数应该被调整为 2^31 − 1
> 4.去掉无用的后导空格

```go
func StrToInt( s string ) int {
    // 去掉无用的前导后导空格
    s = strings.Trim(s," ")
    if len(s) == 0{
        return 0
    }
    // 处理符号位
    sign := 1
    if s[0] == '-'{
        sign = -1
        s = s[1:]
    }else if s[0] == '+'{
        s = s[1:]
    }
    // 判断整数的有效部分
    s1 := []byte(s)
    res := 0
    for _,v := range s1{
        if v>='0' && v<='9'{
            res = res*10 + int(v-'0')*sign
            if sign == 1 && res > math.MaxInt32{
                return math.MaxInt32
            }
            if sign == -1 && res < -math.MaxInt32{
                return -math.MaxInt32-1
            }
        }else{
            break
        }
    }
    return res
}
```

 **JZ20 表示数值的字符串**             

请实现一个函数用来判断字符串str是否表示数值（包括科学计数法的数字，小数和整数）。 

> **科学计数法的数字**(按顺序）可以分成以下几个部分: 
>
>   1.若干空格 
>
>   2.一个整数或者小数 
>
>   3.（可选）一个 'e' 或 'E' ，后面跟着一个整数(可正可负)  
>
>   4.若干空格  
>
> **小数**（按顺序）可以分成以下几个部分：  
>
>   1.若干空格  
>
>   2.（可选）一个符号字符（'+' 或 '-'）  
>
> 3. 可能是以下描述格式之一:  
>
>   3.1 至少一位数字，后面跟着一个点 '.'  
>
>   3.2 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字  
>
>   3.3 一个点 '.' ，后面跟着至少一位数字  
>
>   4.若干空格
>
>  **整数**（按顺序）可以分成以下几个部分：  
>
>   1.若干空格
>   2.（可选）一个符号字符（'+' 或 '-')  
>
>   3. 至少一位数字  
>
>   4.若干空格

```python
class Solution:
    def isNumeric(self , str ):
        # 点的下标
        dotindex = -1
        # 最后一个数字的下表
        numindex = -1
        # e的下标
        eindex = -1
        # 符号下标
        symindex = -1
        li = ["0","1","2","3","4","5","6","7","8","9","+","-",".","e","E"," "]
        # 去掉前置后置空格
        str = str.strip(" ")
        if not str:
            return False
        for i in range(len(str)):
            if str[i] not in li:
                return False
            elif str[i] == "+" or str[i] == "-":
                # 不能符号后没东西
                if i + 1 >= len(str):
                    return False
                # 前面有数字或符号 而且前一个不是e    
                if (numindex != -1 or symindex != -1) and eindex != i-1:
                    return False
                symindex = i
            elif str[i] == "." :
                # 如果前面有e 或有. 或前一个不是数字并且后面没东西 
                if eindex != -1 or dotindex != -1 or ((numindex == -1 or numindex + 1 != i) and i+1 >= len(str) ):
                    return False
                dotindex = i
            elif str[i] == "e" or str[i] == "E":
                # 如果前面没数字或前面有e 或后面没值 
                if numindex == -1 or i+1 >= len(str) or eindex != -1:
                    return False
                eindex = i
            else:
                numindex = i
        return True
```

