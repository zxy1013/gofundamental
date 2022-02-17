**Go语⾔的内存管理与垃圾回收**

**栈分配**，函数调⽤返回后，函数栈帧⾃动销毁(SP下移)，轻量级

**堆分配**，逃逸分析(Escape analysis) ，使用`go build -gcflags="-m" main.go`可以看到逃逸分析的结论，加的-m越多原因输出越详细。



**基本概念-内存管理中的⻆⾊**

内存需要分配	•⾃动 allocator，⼿⼯分配
内存需要回收	•⾃动 collector，⼿⼯回收



**基本概念-内存管理中的三个⻆⾊**

> **Mutator**：其实就是你写的应⽤程序，它会不断地修改对象的引⽤关系，即对象图。
> **Allocator**：内存分配器，负责管理从操作系统中分配出的内存空间，malloc 其实底层就有⼀个内存分配器的实现(glibc 中)，tcmalloc 是 malloc 多线程改进版。Go 中的实现类似tcmalloc。
> **Collector**：垃圾收集器，负责清理死对象，释放内存空间。

![1637550773532](F:\markdown笔记\Go高级工程师实战\image\1637550773532.png)

如果mutator发生逃逸，就必定会调用runtime.newobject函数申请堆空间



**进程虚拟内存布局**

上面是高地址 下面是低地址。栈从高地址向低地址增加，堆从低地址向高地址增加

![1637550986744](F:\markdown笔记\Go高级工程师实战\image\1637550986744.png)

**进程虚拟内存分布 多线程的情况**

主线程从高地址向低地址增加，其余的分布在中间

![1637551296000](F:\markdown笔记\Go高级工程师实战\image\1637551296000.png)



**Allocator基础**

**分类**：

> Bump/Sequential Allocator	不会对释放掉的空闲空间进行复用
>
> Free List Allocator
>
> •First-Fit  每次从链表第一个开始 第一个满足分配条件的，分配所需的 剩余链接
> •Next-Fit 从成功的内存块的下一个节点开始
> •Best-Fit 从头到尾找最佳匹配
> •Segregated-Fit  内存块分为很多级别 8B 16B .. 8K等，每次选择最适合的块的大小，可以减少内存碎片



**malloc实现**

当你执⾏malloc时：

•brk只能通过调整 program break位置推动堆增⻓，小于128KB
•mmap可以从任意未分配位置映射内存，大于等于128KB

**Go 语⾔内存分配**

⽼版本，连续堆，新版本，稀疏堆。申请稀疏堆时，我们该⽤ mmap

> **分配⼤⼩分类：**
> •Tiny : size < 16 bytes && has no pointer(noscan)
> •Small ：has pointer(scan) || (size >= 16 bytes && size <= 32 KB)
> •Large : size > 32 KB

> **内存分配器在 Go 语⾔中维护了⼀个多级结构：**
> mcache -> mcentral -> mheap
> mcache：与 P 绑定，本地内存分配操作，不需要加锁。
> mcentral：中⼼分配缓存，分配时需要上锁，不同 spanClass 使⽤不同的锁。
> mheap：全局唯⼀，从 OS 申请内存，并修改其内存定义结构时，需要加锁，是个全局锁。

**tiny alloc**

在tiny中找 没了去本地队列的5号位置找，没了去mheap分

![M52KFH{4{`QIFU4N12NR9AO](C:\Users\zxy\Documents\Tencent Files\1253141170\FileRecv\MobileFile\Image\8YILVJ]YKC17BOA_71_WE7S.png)

**Small alloc**

![1637552395423](F:\markdown笔记\Go高级工程师实战\image\1637552395423.png)

**Large alloc**

⼤对象分配会直接越过 mcache、mcentral，直接从 mheap 进⾏相应数量的 page 分配。pageAlloc 结构经过多个版本的变化，从：freelist（O(n)）-> treap 二叉树 -> radix tree（比二叉树矮），查找时间复杂度越来越低，结构越来越复杂

> **Refill 流程：**
> •本地 mcache 没有时触发(mcache.refill)
> •从 mcentral ⾥的 non-empty 链表中找(mcentral.cacheSpan)
> •尝试 sweep mcentral 的 empty，insert sweeped -> nonempty(mcentral.cacheSpan)
> •增⻓ mcentral，尝试从 arena 获取内存(mcentral.grow)
> •arena 如果还是没有，向操作系统申请(mheap.alloc)
>
> 最终还是会将申请到的 mspan 放在 mcache 中

![1637553072068](F:\markdown笔记\Go高级工程师实战\image\1637553072068.png)

**堆内存管理-内存分配**
Bitmap 与 allocCache

![1637553258699](F:\markdown笔记\Go高级工程师实战\image\1637553258699.png)

一个mspan有很多elem，分配的元素置为1，分配位和gc标记位



**垃圾回收基础**

> **语义垃圾**(semantic garbage)  — 有的被称作内存泄露
> 语义垃圾指的是从语法上可达(可以通过局部、全局变量引⽤得到)的对象，但从语义上来讲他们是垃圾，垃圾回收器对此⽆能为⼒。slice缩容，后面的元素不释放，依然可以进行底层数组的访问。
> **语法垃圾**(syntactic garbage)
> 语法垃圾是讲那些从语法上⽆法到达的对象，这些才是垃圾收集器主要的收集⽬标。



**常⻅垃圾回收算法**

> **引⽤计数**(Reference Counting)：某个对象的根引⽤计数变为 0 时，其所有⼦节点均需被回收。
>
> **标记压缩**(Mark-Compact)：将存活对象移动到⼀起，解决内存碎⽚问题。
>
> **复制算法**(Copying)：将所有正在使⽤的对象从 From 复制到 To 空间，堆利⽤率只有⼀半。
>
> **标记清扫**(Mark-Sweep)：解决不了内存碎⽚问题。需要与能尽量避免内存碎⽚的分配器使⽤，如 tcmalloc。Go 在这⾥



**Go 语⾔垃圾回收**

**旧流程**

![1637564618127](F:\markdown笔记\Go高级工程师实战\image\1637564618127.png)

**新流程** 缩短第二个stop the word，应用程序阻塞

![1637564782612](F:\markdown笔记\Go高级工程师实战\image\1637564782612.png)

**垃圾回收⼊⼝：**gcStart

触发点：runtime.GC 手动调用 runtime.mallocgc 内存分配 forcegchelper 后台gc触发(2 min)



**GC标记流程**

**GC 标记流程-三⾊抽象**

> ⿊：已经扫描完毕，⼦节点扫描完毕。(gcmarkbits = 1，且在扫描队列外。)
> 灰：已经扫描完毕，⼦节点未扫描完毕。(gcmarkbits = 1, 在扫描队列内)
> ⽩：未扫描，collector 不知道任何相关信息。

1.对象在标记过程中不能丢失
2.Mark 阶段 mutator 的指向堆的指针修改需要被记录下来 (write barrier)
3.GC Mark 的 CPU 控制要努⼒做到 25% 以内



**解决丢失问题的理论基础**

**强三⾊不变性**
strong tricolor invariant 禁⽌⿊⾊对象指向⽩⾊对象
**弱三⾊不变性**
weak tricolor invariant⿊⾊对象指向的⽩⾊对象，如果有灰⾊对象到它的可达路径，那也可以



**垃圾回收代码流程**

> •gcStart ->  gcBgMarkWorker && gcRootPrepare，这时 gcBgMarkWorker 在休眠中 
>
> •schedule -> findRunnableGCWorker 唤醒适宜数量的 gcBgMarkWorker
>
> •gcBgMarkWorker -> gcDrain -> scanobject -> greyobject(set mark bit and put to gcw)
>
> •在 gcBgMarkWorker 中调⽤ gcMarkDone 排空wbBuf 后，使⽤分布式 termination 
> 检查算法，进⼊ gcMarkT ermination -> gcSweep 唤醒后台沉睡的 sweepg 和 scvg -> 
> sweep -> wake bgsweep && bgscavenge

![1637566918092](F:\markdown笔记\Go高级工程师实战\image\1637566918092.png)

