请实现一个函数，将一个字符串s中的每个空格替换成“%20”。 

例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。 

>  """1.统计字符串中空格的个数num,接着在原字符串尾增加长度size=num*2.因为从0变为3
>         2.使用 i 指向扩展前字符串尾部字符，j 指向扩展之后的字符串末尾位置。
>         3.开始遍历字符串
>         a.当s[i]为空格，在s[j]='0',s[j-1]='2',s[j-2]='%';
>         b.当s[i]为字符，s[j]=s[i];"""

```python
# 字符串属于不可变对象，如果需要修改其中的值，只能创建新的字符串对象,原地修改字符串，可以使用io.StringIO对象或array对象
class Solution:
    def replaceSpace(self , s ):
        import io
        count = 0
        prelen = len(s)
        for i in range(prelen):
            if s[i] == " ":
                count += 1
        s = s + ' ' * count * 2
        sio = io.StringIO(s)
        j = len(s) - 1
        for i in range(prelen-1,-1,-1):
            sio.seek(i)
            x = sio.read(1)
            sio.seek(j)
            if x == " ":
                j -= 2
                sio.seek(j)
                sio.write("%20")
                j -= 1
            else
                sio.write(x)
                j -= 1
        return sio.getvalue()      
```

```go
func replaceSpace( s string ) string {
    i := len(s) - 1
    strlist := []byte(s)
    for _,v := range s{
        if v == ' '{
            strlist = append(strlist, ' ')
            strlist = append(strlist, ' ')
        }
    }
    j := len(strlist) - 1
    for i>-1{
        if s[i] == ' '{
            j -= 2
            strlist[j] = '%'
            strlist[j+1] = '2'
            strlist[j+2] = '0'
            j -= 1
        }else{
            strlist[j] = strlist[i]
            j -- 
        }
        i -- 
    }
    return string(strlist)
}
```

