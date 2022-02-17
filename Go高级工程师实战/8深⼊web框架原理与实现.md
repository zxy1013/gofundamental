**middleware实现**

有⼀个 hello world 的业务逻辑，如果想要统计每个请求花的时间，随着业务的迭代，接⼝会逐渐增加，公司发展壮⼤后，我们希望能把耗时可视化，⼏百个接⼝，每个都要改⼀遍？？？

![1638173881097](F:\markdown笔记\Go高级工程师实战\image\1638173881097.png)

![1638173938347](F:\markdown笔记\Go高级工程师实战\image\1638173938347.png)

思路是，把功能性(业务代码)和⾮功能性(⾮业务代码)分离

![1638174036148](F:\markdown笔记\Go高级工程师实战\image\1638174036148.png)

和 HandlerFunc 有相同签名的函数可以强制转换为 HandlerFunc  

HandlerFunc 实现了 ServeHTTP ⽅法，所以 HandlerFunc 实现了 http.Handler 接⼝

**责任链模式、洋葱模式**

![1638174150597](F:\markdown笔记\Go高级工程师实战\image\1638174150597.png)

![1638174166722](F:\markdown笔记\Go高级工程师实战\image\1638174166722.png)

![1638174528702](F:\markdown笔记\Go高级工程师实战\image\1638174528702.png)

**router实现**

![1638174630025](F:\markdown笔记\Go高级工程师实战\image\1638174630025.png)

字典树 trie  设计比较浪费

• 单个节点代表⼀个字⺟ 

• 如果需要对字符串进⾏匹配 

• 只要从根节点开始依次匹配即可

**Radix Tree 前缀树**

同⼀个 URI 在 HTTP 规范中会有多个⽅法，所以每⼀个 method 对应⼀棵 radix tree

![1638156205094](F:\markdown笔记\Go高级工程师实战\image\1638156205094.png)

> node 是 httprouter 树中的节点。 
>
> path  到达节点时，所经过的字符串路径
>
> nType 有⼏种枚举值： 
>
> • static // ⾮根节点的普通字符串节点 
>
> • root // 根节点 
>
> • param(wildcard) // 参数节点，例如 :id 
>
> • catchAll // 通配符节点，例如 *anywayhttprouter 实现细节 

> ⼦节点索引，当⼦节点为⾮参数类型，即本节点的 wildChild 为 false  时，会将每个⼦节点的⾸字⺟放在该索引数组。说是数组，实际上是个string。
>
> 如果⼦节点为参数节点时，indices 应该是个空字符串
>
> wildChild  
>
> 如果⼀个节点的⼦节点中有 param(wildcard) 节点，那么该节点的 wildChild 字段即为 true。 
>
> catchAll 
>
> 以 * 结尾的路由，即为 catchAll。在静态⽂件服务上，catchAll ⽤的⽐较多。后⾯的部分⼀般⽤来描述⽂件路径。如：/software/downloads/monodraw-latest.dmg。

**路由冲突细则** 

> • 在插⼊ wildcard 节点时，⽗节点的 children 数组⾮空且 wildChild 被设置 
>
> 为 false。例如：GET /user/getAll 和 GET /user/:id/getAddr，或者 GET / 
>
> user/*aaa和 GET /user/:id。 
>
> • 在插⼊ wildcard 节点时，⽗节点的 children 数组⾮空且 wildChild 被设置 
>
> 为 true，但该⽗节点的 wildcard ⼦节点要插⼊的 wildcard 名字不⼀样。 
>
> 例如： GET /user/:id/info 和 GET /user/:name/info。 
>
> • 在插⼊ catchAll 节点时，⽗节点的 children ⾮空。例如： GET /src/abc  
>
> 和 `GET /src/*`，或者 `GET /src/:id 和 GET /src/*`。 
>
> • 在插⼊ static 节点时，⽗节点的 wildChild 字段被设置为 true。
>
> •  在插⼊static 节点时，⽗节点的 children ⾮空，且⼦节点 nType 为 catchAll。

**validator实现**

![1638174901620](F:\markdown笔记\Go高级工程师实战\image\1638174901620.png)

看着很恶心

![1638174954412](F:\markdown笔记\Go高级工程师实战\image\1638174954412.png)

也不好

validator改造，深度优先遍历就可以了，有内置校验规则

![1638175013386](F:\markdown笔记\Go高级工程师实战\image\1638175013386.png)

**request binder实现** 简单的⼯⼚模式

不同的⼯⼚实现很简单，就是某种 codec 的 unmarshal

**sql binder实现**

标准库的 API 难⽤且容易犯错，有⽆数的新⼿Gopher倒在sql.Rows忘记关闭的坑下，实现很简单，只要能把特殊开头的单词提取出来就可以了，⽐如:开头，@开头，$开头。

![1638175199122](F:\markdown笔记\Go高级工程师实战\image\1638175199122.png)





