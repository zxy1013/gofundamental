明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了 N 个 1 到 1000 之间的随机整数（ N≤1000 ），对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。请你协助明明完成“去重”与“排序”的工作(同一个测试用例里可能会有多组数据(用于不同的调查)，希望大家能正确处理)。

```python
# 排序 + 去重 先排序后去重
# 快速排序
def fastsort(li):
    if len(li) < 2:
        return li
    # 定义基准值左右两个数列
    mid = li[0]
    left, right = [], []
    for item in li[1:]:
        # 大于基准值放右边
        if item >= mid:
            right.append(item)
        else:
            # 小于基准值放左边
            left.append(item)
    # 使用递归
    return fastsort(left) + [mid] + fastsort(right)

# 去重
def remo(a):
    res = []
    for i in range(1,len(a)):
        if a[i] > a[i-1]:
            res.append(a[i-1])
    res.append(a[-1])
    return res

# 输入使用try以避免EOF when reading a line
def test():
    # 二维数组 存多组随机数
    m = list()
    while True:
       try:
            flag = int(input())
            n = list()
            for j in range(flag):
                k = int(input())
                n.append(k)
            m.append(n)
       except:
           break
    for item in m:
        # 排序
        result = fastsort(item)
        # 去重
        result = remo(result)
        for i in result:
            print(i)
test()
```

