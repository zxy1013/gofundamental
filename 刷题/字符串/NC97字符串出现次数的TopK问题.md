给定一个字符串数组，再给定整数 k ，请返回出现次数前k名的字符串和对应的次数。 返回的答案应该按字符串出现频率由高到低排序。如果不同的字符串有相同出现频率，按字典序排序。 

对于两个字符串，大小关系取决于两个字符串从左到右第一个不同字符的 ASCII 值的大小关系。 比如"ah1x"小于"ahb"，"231"<”32“，字符仅包含数字和字母 

> 堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：
>
> 即子结点的键值或索引总是小于（或者大于）它的父节点。堆排序可以说是一种利用堆的概念来排序的选择排序。根节点索引为0

```python
class Solution:
    def topKstrings(self , strings , k ):
        if k == 0:
            return []
        data = {}
        for i in strings: # 表示str出现count次 str:count
            try:
                data[i] += 1
            except:
                data[i] = 1
        data2 = {}
        for key,val in data.items(): # 表示出现count次的str都有哪些 count:[str1,...]
            try:
                data2[val].append(key)
            except:
                data2[val] = [key]
        res = []
        count = 0
        # 对count排序 从大到小
        keys = sorted(data2.keys(),reverse=True)
        for i in keys:
            # 对str列表进行字典序升序排
            for j in sorted(data2[i]):
                res.append([j,i])
                count += 1
                if count == k:
                    return res
```

