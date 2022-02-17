用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。 

> 输入：["PSH1","PSH2","POP","POP"]
>
> 返回值：1,2
>
> 说明：
>
> "PSH1":代表将1插入队列尾部
> "PSH2":代表将2插入队列尾部
> "POP“:代表删除一个元素，先进先出=>返回1
> "POP“:代表删除一个元素，先进先出=>返回2   

'''
借助栈的先进后出规则模拟实现队列的先进先出
1、当插入时，直接插入 stack1
2、当弹出时，当 stack2 不为空，弹出 stack2 栈顶元素，
如果 stack2 为空，将 stack1 中的全部数逐个出栈，入栈 stack2，再弹出 stack2 栈顶元素
'''

```python
class Solution:
    def __init__(self):
        self.stack1 = []
        self.stack2 = []
    # 后面添加
    def push(self, node):
        self.stack1.append(node)
    # 后面取
    def pop(self):
        # 为空再进
        if not self.stack2:
            while self.stack1:
                self.stack2.append(self.stack1.pop(-1))
        return self.stack2.pop(-1)
```

```go
var stack1 [] int
var stack2 [] int

func Push(node int){
    stack1 = append(stack1, node)
}

func Pop() int{
    if len(stack2) == 0{
        for i := len(stack1)-1 ; i>=0 ; i--{
            stack2 = append(stack2, stack1[i])
        }
        stack1 = []int{}
    }
    temp := stack2[len(stack2)-1]
    stack2 = stack2[:len(stack2)-1]
    return temp
}
```


