**PProf**

https://www.cnblogs.com/gwyy/p/13807267.html

在计算机性能调试领域里，profiling 是指对应用程序的画像，画像就是应用程序使用 CPU 和内存的情况。 Go语言是一个对性能特别看重的语言，因此语言中自带了 profiling 的库，这篇文章就要讲解怎么在 golang 中做 profiling。 想要进行性能优化，首先瞩目在 Go 自身提供的工具链来作为分析依据，本文将带你学习、使用 Go 后花园，涉及如下：

> - runtime/pprof：采集程序（非 Server）的运行数据进行分析
> - net/http/pprof：采集 HTTP Server 的运行时数据进行分析
>

**是什么**

pprof 是用于可视化和分析性能分析数据的工具

**支持什么使用模式**

> - Report generation：报告生成
> - Interactive terminal use：交互式终端使用
> - Web interface：Web 界面
>

**可以做什么**

> - CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置。 报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据 
> - Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏
> - Block Profiling：阻塞分析，记录 goroutine 阻塞等待同步（包括定时器通道）的位置。 报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈 
> - Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况
> - Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的 
>

**一个简单的例子**

**编写 demo 文件**

（1）demo.go，文件内容：

```go
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"pproftest/data"
)

func main() {
	go func() {
		for {
            log.Println(data.Add("hello test pprof usage"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
```

（2）data/d.go，文件内容：

```go
package data

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
```

运行这个文件，你的 HTTP 服务会多出 /debug/pprof 的 endpoint 可用于观察应用程序的情况

pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取各个函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，形成一个性能分析报告。 

**为什么要初始化 net/http/pprof**

在前面例子中，在引用上对 net/http/pprof 包进行了默认的初始化，（也就是 _） ，如果没在对该包进行初始化，则无法调用pprof的相关接口，这是为什么呢，我们可以一起看看该包初始化方法：

```go
func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
	http.HandleFunc("/debug/pprof/trace", Trace)
}
```

实际上，在初始化函数中， net/http/pprof 会对标准库中的 net/http 默认提供 DefaultServeMux 进行路由注册

我们在例子中使用的 HTTP Server,也是标准库默认提供的，因此便可以注册进去。

在实际项目中 我们有独立的 ServeMux的，这时候只需要将PProf对应的路由注册进去即可　　

**分析**

**一、通过 Web 界面**

查看当前总览：访问 `http://127.0.0.1:6060/debug/pprof/`

```gradle
/debug/pprof/

profiles:
0    block
5    goroutine
3    heap
0    mutex
9    threadcreate

full goroutine stack dump
```

这个页面中有许多子页面，咱们继续深究下去，看看可以得到什么？

> - cpu（CPU Profiling）: `$HOST/debug/pprof/profile`，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
> - block（Block Profiling）：`$HOST/debug/pprof/block`，查看导致阻塞同步的堆栈跟踪
> - goroutine：`$HOST/debug/pprof/goroutine`，查看当前所有运行的 goroutines 堆栈跟踪
> - heap（Memory Profiling）: `$HOST/debug/pprof/heap`，查看活动对象的内存分配情况
> - mutex（Mutex Profiling）：`$HOST/debug/pprof/mutex`，查看导致互斥锁的竞争持有者的堆栈跟踪
> - threadcreate：`$HOST/debug/pprof/threadcreate`，查看创建新OS线程的堆栈跟踪
>

![1639216119312](F:\markdown笔记\Go高级工程师实战\image\1639216119312.png)

**二、通过交互式终端使用**

 通过命令行完整对正在运行的程序 pprof进行抓取和分析 

```bash
PS E:\gopro\src\fundation\ppoftest> go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60

Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile?seconds=60
Saved profile in C:\Users\zxy\pprof\pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Dec 11, 2021 at 4:10pm (CST)
Duration: 60.22s, Total samples = 35.74s (59.35%)
Entering interactive mode (type "help" for commands, "o" for options)
```

执行该命令后，需等待60秒（可调整 seconds 的值），pprof 会进行 CPU Profiling。结束后将默认进入 pprof 的交互式命令模式，可以对分析的结果进行查看或导出。 输入命令 top 10 查看对应资源开销 （例如 cpu就是执行耗时/开销 Memory 就是内存占用大小）排名前十的函数 

```bash
(pprof) top 10
Showing nodes accounting for 31.20s, 87.30% of 35.74s total
Dropped 157 nodes (cum <= 0.18s)
Showing top 10 nodes out of 45
      flat  flat%   sum%        cum   cum%
    23.66s 66.20% 66.20%     23.92s 66.93%  runtime.cgocall
     2.70s  7.55% 73.75%      2.70s  7.55%  runtime.memmove
     2.07s  5.79% 79.55%      2.07s  5.79%  runtime.memclrNoHeapPointers
     1.10s  3.08% 82.62%      1.77s  4.95%  runtime.scanobject
     0.47s  1.32% 83.94%      0.47s  1.32%  runtime.procyield
     0.38s  1.06% 85.00%      2.99s  8.37%  runtime.mallocgc
     0.22s  0.62% 85.62%      0.29s  0.81%  log.itoa
     0.22s  0.62% 86.23%      0.22s  0.62%  time.absDate
     0.19s  0.53% 86.77%     26.14s 73.14%  log.(*Logger).Output
     0.19s  0.53% 87.30%      0.87s  2.43%  log.(*Logger).formatHeader
```

- flat：当前函数上运行耗时

- flat%：函数自身占用的 CPU 运行耗时总比例

- sum%：函数自身累积使用 CPU 总比例

- cum：当前函数及其调用函数的运行总耗时

- cum%：函数自身及其调用函数占 CPU 运行耗时总比例

- 最后一列为函数名称,在大多数的情况下，我们可以通过这五列得出一个应用程序的运行情况，加以优化

  **Heap Profiling** 

```bash
PS E:\gopro\src\fundation\ppoftest> go tool pprof http://localhost:6060/debug/pprof/heap
Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
Saved profile in C:\Users\zxy\pprof\pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
Type: inuse_space
Time: Dec 11, 2021 at 4:22pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)

(pprof) top
Showing nodes accounting for 29537.10kB, 100% of 29537.10kB total
Showing top 10 nodes out of 28
      flat  flat%   sum%        cum   cum%
27488.38kB 93.06% 93.06% 27488.38kB 93.06%  pproftest/data.Add (inline)
  512.50kB  1.74% 94.80%   512.50kB  1.74%  runtime.allocm
  512.20kB  1.73% 96.53%   512.20kB  1.73%  runtime.malg
  512.02kB  1.73% 98.27%   512.02kB  1.73%  syscall.(*DLL).FindProc
     512kB  1.73%   100%      512kB  1.73%  runtime.doaddtimer
         0     0%   100%   512.02kB  1.73%  internal/poll.checkSetFileCompletionNotificationModes
         0     0%   100%   512.02kB  1.73%  internal/poll.init.0
         0     0%   100% 27488.38kB 93.06%  main.main.func1
         0     0%   100%      512kB  1.73%  runtime.bgscavenge
         0     0%   100%   512.02kB  1.73%  runtime.doInit

```

- -inuse_space：分析应用程序的常驻内存占用情况
- -alloc_objects：分析应用程序的内存临时分配情况
-  inuse_objects 和 alloc_space 类别，分别对应查看每个函数的对象数量和分配的内存空间大小。 

**三、PProf 可视化界面**

这是令人期待的一小节。在这之前，我们需要简单的编写好测试用例来跑一下

##### 编写测试用例

（1）新建 data/d_test.go，文件内容：

```go
package data
import "testing"
const url = "hello test pprof usage"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
```

（2）执行测试用例

```bash
PS E:\gopro\src\fundation\ppoftest\data> go test -bench=Add -cpuprofile=cpu
ok      pproftest/data  0.402s
```

##### 启动 PProf 可视化界面

需要安装 graphviz，[参考：Graphviz安装及简单使用](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.cnblogs.com%2Fshuodehaoa%2Fp%2F8667045.html)。 

方法一：

```routeros
PS E:\gopro\src\fundation\ppoftest\data> go tool pprof -http=:8080 cpu
```

##### 查看 PProf 可视化界面 

可以在http://localhost:8080/ui 的view中切换查看视图

通过 PProf 的可视化界面，我们能够更方便、更直观的看到 Go 应用程序的调用链、使用情况等，并且在 View 菜单栏中，还支持如上多种方式的切换

或者安装

```vim
$ go get -u github.com/google/pprof
```

它就是本次的目标之一，它的最大优点是动态的。调用顺序由上到下（A -> B -> C -> D），每一块代表一个函数，越大代表占用 CPU 的时间更长。同时它也支持点击块深入进行分析！

理论上等价实际上相差几个数量级 string是指针类型，数组值类型

![1639215481261](F:\markdown笔记\Go高级工程师实战\image\1639215481261.png)

