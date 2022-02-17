**信息源**：Github Trending、reddit、medium、hacker news，morning paper，acm.org，oreily，国外的领域相关⼤会(如 OSDI，SOSP，VLDB)论⽂。

> 为什么 Go 语⾔适合现代的后端编程环境？
> • 服务类应⽤以 API 居多，IO 密集型，且⽹络 IO 最多
> • 运⾏成本低，⽆ VM虚拟机。⽹络连接数不多的情况下内存占⽤低。
> • 强类型语⾔，易上⼿，易维护。

**对 Go 的启动和执⾏流程建⽴简单的宏观认识**

**理解可执⾏⽂件**

打开VM上的虚拟机，简单断点

```cmd
systemctl start docker
sudo docker pull golang
sudo docker run -it --name go-demo golang bash // 用镜像golang创建了一个名为go-demo的容器，在容器中创建了一个 Bash会话。
docker ps -a --no-trunc  // 查看镜像id
sudo docker start 1b96fc3c9eac0fa5e37064cc2b1a0ab14950ac99f7415ab60e50d9f2895 // 开启镜像
sudo docker exec -it 1b96fc3c9eac0fa5e37064cc2b1a0ab14950ac99f7415ab60e50d9f28956748d /bin/bash // 进入
// 安装需要的包
apt-get update // 升级
go env -w GOPROXY=https://goproxy.cn
go get -u github.com/go-delve/delve/cmd/dlv
apt-get install binutils -y
apt-get install vim -y
apt-get install gdb -y

// 创建文件
cd src
touch hello.go
vi hello.go
go mod init gotest
go build hello.go
go run hello.go


'''
    使用b命令打断点，有三种方法：
    b + 地址
    b + 代码行数
    b **.go:5
    b + 函数名
'''
readelf -h ./hello // 找Entry point address
dlv exec ./hello // 开启断点调试状态
b *0x45c220 // 在entry point处打断点
c // 执行到下一个断点
si // 一次执行一条汇编指令
exit // 退出调试

dlv exec ./hello // 断点
b *0x45c220
si
disass // 反汇编
```

查看go语言的runtime的下列函数

```cmd
dlv exec ./hello // 断点
b *0x45c220
si // 执行到JMP再si进入
n // go语言内部可以用n表示下一行
b *runqput // 在runqput处打断点
c // 进入下一个断点 - runqput
stack或bk  // 查看函数从哪里跳转进来
frame [0-n] // 可以查看上一步列出的具体函数

r // 重新开始调试
c
c
dlv exec ./hello
b runqput // 在runqput处打断点
b chansend // 在chansend处打断点
exit
```

查看向关闭chan发送数据具体报错函数的位置

```cmd
vim sendclosedchan.go
package main
func main(){
	var ch = make(chan int)
	close(ch)
	ch <-1
}
dlv debug sendclosedchan.go
b main.main // main函数打断点
c // 进入main函数
n  // 进入下一行 3次
disass // 查看报错的此行代码都是在干啥 看到call打断点
b runtime.chansend1 // 报错函数打断点
c // 进入报错函数
s 
s // go代码跳入函数内部,一直n找到需要的语句
```

**Go 程序 hello.go 的编译过程：**

⽂本 --> 编译，通过`go build **.go` 生成⼆进制可执⾏⽂件

使用`go build -x .go` 可以详细观察编译过程：**编译** (将文本代码转为目标文件)加 **链接**(将编译出来的目标文件和已经编译好的静态标准库文件做链接) 后将编译好的完整的文件输出到临时目录，最后挪动到原目录后删除。

可执⾏⽂件在不同的操作系统上规范不⼀样，Linux 的可执⾏⽂件 ELF(Executable and Linkable Format) 由三部分构成：ELF header	Section header	Sections

操作系统执⾏可执⾏⽂件的步骤(以 linux 为例)：
解析ELF Header，通过Header的内容将段加载⽂件内容⾄内存，从 entry point 开始执⾏代码

通过 entry point 找到 Go 进程的执⾏⼊⼝，使⽤` readelf -h ./x` ，在 dlv 调试器中 `b *entry_addr` 找到代码位置

**Go 进程的启动与初始化**

计算机是怎么执⾏我们的程序的呢？
CPU ⽆法理解⽂本，只能执⾏⼀条⼀条的⼆进制机器码指令，每次执⾏完⼀条指令，pc 寄存器就指向下⼀条继续执⾏，在 64 位平台上 pc 寄存器 = rip

Go 语⾔是⼀⻔有 runtime 的语⾔，那么 runtime 是什么？
可以认为 runtime 是为了实现额外的功能，⽽在程序运⾏时⾃动加载/运⾏的⼀些模块。

> **Go 语⾔的 runtime 包括：**
> Scheduler	调度器管理所有的 G，M，P，在后台执⾏调度循环
> Netpoll	⽹络轮询负责管理⽹络 FD 相关的读写、就绪事件
> Memory	当代码需要内存时，负责内存分配⼯作	
> Garbage	当内存不再需要时，负责回收内存
> 这些模块中，最核⼼的就是 Scheduler，它负责串联所有的 runtime 流程

通过 entry point 找到 Go 进程的执⾏⼊⼝
找到runtime.rt0_go -- > argc/argv处理 -- > 全局m0/g0初始化(m0: Go 程序启动后创建的第⼀个线程) -- > 获取CPU核⼼数 -- > 初始化内置数据结构 -- > 开始执⾏⽤户main函数(开始进⼊调度循环)

**调度组件与调度循环**

Go 的调度流程本质上是⼀个⽣产(`go func(){}`) -- 消费(线程M)流程，用户提交计算任务到队列(协程G)。

![1636619093707](F:\markdown笔记\Go高级工程师实战\image\1636619093707.png)

**goroutine 的⽣产端**：多级队列减少锁竞争

![1636547359187](F:\markdown笔记\Go高级工程师实战\image\1636547359187.png)

本地队列的前一半和踢走的老g组成一个新的结构batch，拼成一个链表，存入全局链表中。本地队列全局链表的原因：本地局部性数量有限256，数组访问速度快；全局无大小限制，需要链表。理论上新增的g比老的g优先级高，因为局部性原理。

```go
func main() {
	runtime.GOMAXPROCS(1)
	for i:=0;i<10;i++{
		i:= i
		go func(){
			fmt.Printf("A:%d ",i)
		}()
	}
	var ch = make(chan int)
	<- ch
}
func main() {
	runtime.GOMAXPROCS(1)
	for i:=0;i<10;i++{
		i:= i
		go func(){
			fmt.Printf("A:%d ",i)
		}()
	}
	time.Sleep(time.Second)
}
// 结果相同 旧版本的time.Sleep会开一个新的goroutine输出0-9，新版本的不会，所以先输出9，因为只有一个p，以及新增的g比老的g优先级高
// A:9 A:0 A:1 A:2 A:3 A:4 A:5 A:6 A:7 A:8
```

**goroutine 的消费端**：M 执⾏调度循环时，必须与⼀个 P (逻辑概念 绑定了一堆队列和缓存)绑定，调度循环实际上就是 Go 程序在启动的时候，会创建和 CPU 核心数相等个数的 P，会创建初始的 m，称为 m0。这个 m0 会启动一个调度循环：不断地找 g，执行，再找 g，随着程序的运行，m 更多地被创建出来，因此会有更多的调度循环在执行。那边生产者在不断地生产 g，这边 m 的调度循环不断地在消费 g，整个过程就 run 起来了。找 g 的过程中当然也是从三级队列里找。

![1636548181875](F:\markdown笔记\Go高级工程师实战\image\1636548181875.png)

每60次访问一次全局链表(每次运行到execute计数器加1  当%61 == 0，拿取全局链表中第一个G在线程中执行)，平常访问local queue(先判断runnext，以及local的第一个)。当local为空，去全局找，全局非空，拿取前一半(上限为128个)并取出头部元素执行，后面的元素塞入local queue，若全局空，去其他的p中找，从尾部偷取其他的local queue的一半，执行最后一个元素，将其他的存入local queue。Work stealing 就是说的 runqsteal -> runqgrab 这个流程。

> 现在我们再来看看这些⽂字定义
> G：goroutine，⼀个计算任务。由需要执⾏的代码和其上下⽂组成，上下⽂包括：当前代码位置，栈顶、栈底地址，状态等。
> M：machine，系统线程，执⾏实体，想要在 CPU 上执⾏代码，必须有线程，与 C 语⾔中的线程相同，通过系统调⽤ clone 来创建。
> P：processor，虚拟处理器，M 必须获得 P 才能执⾏代码，否则必须陷⼊休眠(后台监控线程除外)，你也可以将其理解为⼀种 token，有这个 token，才有在物理 CPU 核⼼上执⾏的权⼒。

**处理阻塞**

无缓冲区channel -- time.sleep -- 网络读写 -- 锁操作

这些情况不会阻塞调度循环，⽽是会把 goroutine 挂起，所谓的挂起，其实让 g 先进入某个数据结构，待 ready 后再继续执⾏，不会占⽤线程，这时候，线程会进⼊ schedule，继续消费队列，执⾏其它的 g。

为啥有的等待是 sudog，有的是 g?
⼀个 g 可能对应多个 sudog，⽐如⼀个 g 会同时 select 多个channel，所以在同步状态下需要将g封装成sudog。

前⾯这些都是能被 runtime 拦截到的阻塞，还有⼀些是 runtime **⽆法拦截**的：
sys	CGO，在执⾏ c 代码，或者阻塞在 syscall 上时，必须占⽤⼀个线程。⻓时间运⾏需要剥离 P 执⾏。

**处理阻塞：**
sysmon: system monitor
⾼优先级，在专有线程中执⾏，不需要绑定 P 就可以执⾏，作用：

> checkdead	检查是否当前所有的线程都被阻塞，如果是，则崩溃
> netpoll 	网络轮询 inject g list to global run queue	
> retake	如果是 syscall 卡了很久10ms，那就把 p 剥离(handoff p) ，等待后面主动恢复/ 如果是⽤户 g 运⾏很久10ms，那么发信号抢占，那就把 p 剥离(handoff p) 将其放入全局队列
>
> ```go
> func main() {
> 	var i =1
> 	go func(){
> 		for {
> 			i++
> 		}
> 	}()
> }
> // GC时需要停⽌所有goroutine,⽽⽼版本的Go的g停⽌需要主动让出,1.14增加基于信号的抢占之后，该问题被解决,解决了死循环导致进程hang死问题
> ```

**调度器的发展历史**

![1636552004355](F:\markdown笔记\Go高级工程师实战\image\1636552004355.png)

Goroutine ⽐ Thread 优势在哪？
![1636552280692](F:\markdown笔记\Go高级工程师实战\image\1636552280692.png)

**goroutine 的切换成本**
gobuf 描述⼀个 goroutine 所有现场，从⼀个 g 切换到另⼀个 g，只要把这⼏个现场字段保存下来，再把 g 往队列⾥⼀扔，m 就可以执⾏其它 g 了，⽆需进⼊内核态。

各种阻塞场景是怎么在代码⾥找到的？要知道 runtime 中可以接管的阻塞是通过 gopark/goparkunlock 挂起和 goready 恢复的，那么我们只要找到 runtime.gopark 的调⽤⽅，就可以知道在哪些地⽅会被 runtime 接管了。即使⽤ IDE 查看函数的调⽤⽅。

Goland 点击查找选项输入需要查找的函数gopark，点击第一个进入文件proc.go,点击函数名右键选择find usages，Scope中选择Project and Libraries即可得到调用方。

