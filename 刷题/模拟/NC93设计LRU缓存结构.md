设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为 k ，并有如下两个功能 

  \1. set(key, value)：将记录(key, value)插入该结构 

  \2. get(key)：返回key对应的value值 

提示: 

> 1.某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的，然后都会刷新缓存。 
>
> 2.当缓存的大小超过k时，移除最不经常使用的记录。 
>
> 3.输入一个二维数组与k，二维数组每一维有2个或者3个数字，第1个数字为opt，第2，3个数字为key，value  
>
> 若opt=1，接下来两个整数key, value，表示set(key, value)
>
> 若opt=2，接下来一个整数key，表示get(key)，若key未出现过或已被移除，则返回-1 对于每个opt=2，输出一个答案
>
> 4.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹 



> ''' 如果想让字典有序，可以使用collections.OrderedDict
> from collections import OrderedDict
> OrderedDict是记住键首次插入顺序的字典。如果新条目覆盖现有条目，则原始插入位置保持不变。
> 删除条目并重新插入会将其移动到末尾。
> '''

```python
from collections import OrderedDict
class Solution:
    def LRU(self , operators , k ):
        # 创建结果列表
        res = []
        # 创建有序数组d
        d = OrderedDict()
        for op in operators:
            if op[0] == 1:
                ans = d.get(op[1],-1)
                # 需要添加新元素
                if ans == -1:
                    if len(d) == k:
                        '''popitem()方法会删除并返回(key, value)对。
                        如果last为真，则以LIFO(后进先出)顺序返回这些键值对，
                        如果为假，则以FIFO(先进先出)顺序返回。'''
                        d.popitem(last=False)
                d[op[1]] = op[2]
                # 更新位置
                d.move_to_end(op[1])
            if op[0] == 2:
                '''move_to_end(key, last=True),该方法用于将一个已存在的key移动到有序字典的任一端。
                如果last为True（默认值），则移动到末尾，如果last为False，则移动到开头。
                如果key不存在，引发KeyError'''
                ans = d.get(op[1],-1)
                res.append(ans)
                if ans != -1:
                    d.move_to_end(op[1])
        return res
```



```go
package main

func deletel(a[]int,k int)[]int{
    res := make([]int,0)
    for _,v := range a{
        if v != k{
            res = append(res,v)
        }
    }
    return res
}
func LRU( operator [][]int ,  k int ) []int {
    res := make([]int,0)
    // slice 存key的位置前后信息 + map 存key value
    location := make([]int,0)
    dict := make(map[int]int)
    for i,_ := range operator{
        if operator[i][0] == 1{ // 插入
            // 查看缓存中是否有这个数
            _,ok := dict[operator[i][1]]
            // 缓存满了 先删除后添加
            if !ok{
                if len(location) == k{
                    // 删除缓存
                    delete(dict,location[0])
                    // 删除位置信息
                    location = deletel(location,location[0])
                }else{
                    location = deletel(location, operator[i][1])
                }
            }
            // 添加
            dict[operator[i][1]] = operator[i][2]
            location = append(location, operator[i][1])
        }else{ // 查询
            v,ok := dict[operator[i][1]]
            if ok{ // 有值 查询后更新位置信息
                res = append(res,v)
                // 删除原来的下标，追加在最后
                location = deletel(location,operator[i][1])
                location = append(location,operator[i][1])
            }else{
                res = append(res,-1)
            }
        }
    }
    return res
}
```

 **NC94 设计LFU缓存结构**  

一个缓存结构需要实现如下功能。 

-    set(key, value)：将记录(key, value)插入该结构    
-    get(key)：返回key对应的value值   

但是缓存结构中最多放K条记录，如果新的第K+1条记录要加入，就需要根据策略删掉一条记录，然后才能把新记录加入。这个策略为：在缓存结构的K条记录中，哪一个key从进入缓存结构的时刻开始，被调用set或者get的次数最少，就删掉这个key的记录。如果调用次数最少的key有多个，上次调用发生最早的key被删除， 这就是LFU缓存替换算法。实现这个结构，K作为参数给出

```python
class Solution:
    def LFU(self , operators , ki ):
        res = []
        from collections import OrderedDict
        # 创建有序数组d 记录顺序 键和值和次数 {key:[val,count]}
        d = OrderedDict()
        for items in operators:
            if items[0] == 1:
                ans = d.get(items[1],-1)
                # 需要添加新元素
                if ans == -1:
                    # 删除记录
                    if len(d) == ki:
                        keys = list(d.keys())
                        min_key = keys[0]
                        for i in range(1,len(keys)):
                            if d[keys[i]][1] < d[min_key][1]:
                                min_key = keys[i]
                        del d[min_key]
                # 更新次数和位置
                d[items[1]] = [items[2],d.get(items[1],0)+1]
                d.move_to_end(items[1])
            if items[0] == 2:
                ans = d.get(items[1],[])
                if ans != []:
                    res.append(ans[0])
                    # 更新调用次数
                    d[items[1]][1] += 1
                    # 更新调用顺序
                    d.move_to_end(items[1],last=True)
                else:
                    res.append(-1)
        return res
```

