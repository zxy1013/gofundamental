假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益 

1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天 

2.如果不能获取到任何利润，请返回0 

3.假设买入卖出均无手续费

双指针解决 找最大差值

>         '''一个指针记录访问过的最小值（注意这里是访问过的最小值），一个指针一直往后走，
>         然后计算他们的差值，保存最大的即可'''

```python
class Solution:
    def maxProfit(self , prices):
        mini = prices[0]
        result = 0
        for i in range(len(prices)):
            mini = min(mini,prices[i])
            result = max(result,prices[i]-mini)
        return result
```

```go
func maxProfit( prices []int ) int {
    min := prices[0]
    res := 0
    for i,_ := range prices{
        if prices[i] < min{
            min = prices[i]
        }
        if prices[i] - min > res{
            res = prices[i] - min
        }
    }
    return res
}
```

 **NC134 买卖股票的最好时机(二)**             

假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益 

  \1. 你可以多次买卖该只股票，但是再次购买前必须卖出之前的股票 

  \2. 如果不能获取收益，请返回0 

  \3. 假设买入卖出均无手续费 

```python
class Solution:
    def maxProfit(self , prices ):
        if not prices:
            return 0
        if sorted(prices) == prices:
            return prices[-1] - prices[0]
        if sorted(prices,reverse=True) == prices:
            return 0
        sumi = 0
        # 只要涨一天就可以买入卖出
        for i in range(len(prices)-1):
            if prices[i+1] - prices[i] > 0:
                sumi += prices[i+1] - prices[i]
        return sumi
```

```go
func maxProfit( prices []int ) int {
    if len(prices) == 0{
        return 0
    }
    res := 0
    // 因为每天都可以卖出，所以只要大于就可以累加
    for i,_ := range prices[1:]{
        if prices[i+1] > prices[i]{
            res += (prices[i+1] - prices[i])
        }
    }
    return res
}
```

 假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
 \1. 你最多可以对该股票有k笔交易操作，一笔交易代表着一次买入与一次卖出，但是再次购买前必须卖出之前的股票
 \2. 如果不能获取收益，请返回0
 \3. 假设买入卖出均无手续费 

>   根据递归的逻辑，我们可以改出动态规划。   
>
> 1. ​    分析递归的参数列表：可变参数一共有rest剩余几次操作，pocess 0和1表示未拥有和拥有 ，因此本质上是一个2维动态规划问题。      
> 2. ​    根据参数的取值范围，确定dp表的维度：rest的取值范围是[0,k]，pocess 为布尔类型只有两个取值，这里把它规定为0和1。因此dp表的维度可以规定为`(k+1)*2`。
> 3. ​    base case：某天第一天拥有股票，则收益为股价的相反数，即`dp[x][1] = -prices[0]`。      因为k次交易，如果有一次交易 那么赋值的范围为[0,k) 表示剩余几次交易次数 以及手里拥有股票的最大收益
> 4.    要保证第day天手上无股票，可以在前一天持股的基础上卖掉也可以直接从前一天无股的状态转移过来
>       要保证第day天手上有股票，可以在前一天无股的基础上买入也可以直接从前一天有股的状态转移过来      

```go
package main

func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}
func maxProfit( prices []int ,  k int ) int {
    if k == 0 || len(prices) < 2{
        return 0
    }
    // 记录剩余买卖的次数以及手上是否拥有股票
    dp := make([][]int, k+1)
    // 0 1 代表未拥有 拥有
    for i,_ := range dp{
        dp[i] = make([]int,2)
    }
    // 首日买入 不论剩余几次买卖次数，均为负债情况
    // 剩余i次交易次数 以及手里拥有股票的最大收益
    for i := 0;i<k;i++{
        dp[i][1] = -prices[0]
    }
    // 第一天为初始 从第二天开始
    for day := 1; day < len(prices);day ++{
        // 剩余交易次数
        for rest := k-1;rest>-1;rest--{
            // 手上无股票，可以在前一天持股的基础上卖掉
            // 也可以直接从前一天无股的状态转移过来
            dp[rest][0] = max(dp[rest][1]+prices[day],dp[rest][0])
            // 手上有股票，可以在前一天无股的基础上买入
            // 也可以直接从前一天有股的状态转移过来
            // 开启新 买入，次数减一
            // 因为次数增加的时候对应的是未拥有状态，所以不需要考虑k的特殊情况
            dp[rest][1] = max(dp[rest+1][0]-prices[day],dp[rest][1])
        }
    }
    // 剩余0次 手上无股票
    return dp[0][0]
}
```

 假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
 \1. 你最多可以对该股票有两笔交易操作，一笔交易代表着一次买入与一次卖出，但是再次购买前必须卖出之前的股票
 \2. 如果不能获取收益，请返回0
 \3. 假设买入卖出均无手续费 

```go
package main

func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}
func maxProfit( prices []int) int {
    if len(prices) < 2{
        return 0
    }
    // 记录剩余买卖的次数以及手上是否拥有股票
    dp := make([][]int, 3)
    // 0 1 代表未拥有 拥有
    for i,_ := range dp{
        dp[i] = make([]int,2)
    }
    // 首日买入 不论剩余几次买卖次数，均为负债情况
    // 剩余i次交易次数 以及手里拥有股票的最大收益
    for i := 0;i<2;i++{
        dp[i][1] = -prices[0]
    }
    // 第一天为初始 从第二天开始
    for day := 1; day < len(prices);day ++{
        // 剩余交易次数
        for rest := 2-1;rest>-1;rest--{
            // 手上无股票，可以在前一天持股的基础上卖掉
            // 也可以直接从前一天无股的状态转移过来
            dp[rest][0] = max(dp[rest][1]+prices[day],dp[rest][0])
            // 手上有股票，可以在前一天无股的基础上买入
            // 也可以直接从前一天有股的状态转移过来
            // 开启新 买入，次数减一
            // 因为次数增加的时候对应的是未拥有状态，所以不需要考虑k的特殊情况
            dp[rest][1] = max(dp[rest+1][0]-prices[day],dp[rest][1])
        }
    }
    // 剩余0次 手上无股票
    return dp[0][0]
}
```



