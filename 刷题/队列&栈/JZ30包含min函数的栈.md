定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。  

> 此栈包含的方法有： 
>
> push(value):将value压入栈中 
>
> pop():弹出栈顶元素 
>
> top():获取栈顶元素 
>
> min():获取栈中最小元素 
>
> 示例: 
>
> 输入:  ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"] 
>
> 输出:  -1,2,1,-1 
>
> "PSH-1"表示将-1压入栈中，栈中元素为-1  
>
> "PSH2"表示将2压入栈中，栈中元素为2，-1  
>
> “MIN”表示获取此时栈中最小元素==>返回-1 
>
> "TOP"表示获取栈顶元素==>返回2  
>
> "POP"表示弹出栈顶元素，弹出2，栈中元素为-1 
>
> "PSH1"表示将1压入栈中，栈中元素为1，-1
>
> "TOP"表示获取栈顶元素==>返回1   
>
> “MIN”表示获取此时栈中最小元素==>返回-1

'''
使用辅助栈
需要一个正常栈normal,用于栈的正常操作,然后需要一个辅助栈minval,专门用于获取最小值
4 2 3 1
normal 4 2 3 1
minval 4 2 2 1
'''

```python
class Solution:
    def __init__(self):
        self.stack = []
        self.stack_min = []
    def push(self, node):
        self.stack.append(node)
        if self.stack_min and self.stack_min[-1] < node:
            self.stack_min.append(self.stack_min[-1])
        else:
            self.stack_min.append(node)
    def pop(self):
        self.stack_min.pop(-1)
        self.stack.pop(-1)
    def top(self):
        if self.stack:
            return self.stack[-1]
    def min(self):
        if self.stack_min:
            return self.stack_min[-1]
```

```go
var stack []int
var stack_min []int
func Push(node int) {
    stack = append(stack,node)
    if len(stack_min) > 0 && node > stack_min[len(stack_min)-1]{
        stack_min = append(stack_min,stack_min[len(stack_min)-1])
    }else{
        stack_min = append(stack_min,node)
    }
}
func Pop() {
    stack = stack[:len(stack)-1]
    stack_min = stack_min[:len(stack_min)-1]
}
func Top() int {
    return stack[len(stack)-1]
}
func Min() int {
    return stack_min[len(stack_min)-1]
}
```

