**Go并发编程最佳实践**

**并发内置数据结构**

> **sync.Once** 只有⼀个⽅法，Do() 但 o.Do 需要保证： 
>
> • 初始化⽅法必须且只能被调⽤⼀次 
>
> • Do 返回后，初始化⼀定已经执⾏完成

> **sync.Pool**  
>
> 主要在两种场景使⽤： 
>
> • 进程中的 `inuse_objects` 数过多，gc mark 消耗⼤量 CPU 
>
> • 进程中的`inuse_objects` 数过多，进程 RSS 占⽤过⾼ 
>
> 请求⽣命周期开始时，pool.Get，请求结束时，pool.Put。 
>
> 在 fasthttp 中有⼤量应⽤

> **semaphore**  
>
> 是锁的实现基础， 所有同步原语的基础设施
>
> ![1637672122216](F:\markdown笔记\Go高级工程师实战\image\1637672122216.png)

> **sync.Mutex**
>
> ![1637672167116](F:\markdown笔记\Go高级工程师实战\image\1637672167116.png)
>
> **sync.RWMutex**
>
> ![1637672197690](F:\markdown笔记\Go高级工程师实战\image\1637672197690.png)

> **sync.Waitgroup** 
>
> • Counter 减到 0 时，要唤醒所有 sema 上阻塞的 sudog
>
> ![1637672548052](F:\markdown笔记\Go高级工程师实战\image\1637672548052.png)

**常⻅并发bug**

> 1 死锁
>
> 2 Map concurrent writes/reads  Map的并发读写
>
> 3 Channel 关闭 panic  
>
> 1. M receivers, one sender, the sender says "no more sends" by  
>
> closing the data channel  
>
> 2. One receiver, N senders, the only receiver says "please stop  
>
> sending more" by closing an additional signal channel  
>
> 3. M receivers, N senders, any one of them says "let's end the  
>
> game" by notifying a moderator to close an additional signal  
>
> channel

> **Happen-before** 
>
> 初始化： 
>
> 1 A pkg import B pkg，那么 B pkg 的 init 函数⼀定在 A pkg 的 init 函数之前执⾏。 Init 函数⼀定在 main.main 之前执⾏内存模型 
>
> 2 Goroutine 创建 Goroutine 的创建(creation)⼀定先于 goroutine 的执⾏ (execution) 
>
> 3 Channel 收/发：A send on a channel happens before the corresponding receive from that channel completes. 
>
> 4 Channel 收/发：The closing of a channel happens before a receive that returns a zero value because the channel is  closed. 
>
> 5 Channel 收/发：A receive from an unbuffered channel happens before the send on that channel completes. ⽆ buffer 的 chan receive 先于 send 执⾏完
>
> 6 Lock：For any sync.Mutex or sync.RWMutex variable l and  n < m, call n of l.Unlock() happens before call m of l.Lock()  returns.  Unlock ⼀定先于 Lock 函数返回前执⾏完
>
> 7 Once：A single call of f() from once.Do(f) happens (returns)  before any call of once.Do(f) returns
>
> 本质是在⽤户不知道 memory barrier 概念和具体实现的前提 下，能够按照官⽅提供的 happen-before 正确进⾏并发编程。