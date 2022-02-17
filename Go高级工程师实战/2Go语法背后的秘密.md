![1636632744174](F:\markdown笔记\Go高级工程师实战\image\1636632744174.png)

**中间代码(SSA)⽣成与优化 **https://golang.design/gossa

SSA(Single Static Assignment)的两⼤要点是： 

Static: 每个变量只能赋值⼀次(因此应该叫常量更合适)；

Single: 每个表达式只能做⼀个简单运算，对于复杂的表达式ab+cd要拆分成: t0=ab; t1=cd; t2=t0+t1; 三个简单表达式；

机器码生成网址 https://godbolt.org/

**编译原理基础**

编译后所有函数地址都是从0开始，每条指令是相对函数第⼀条指令的偏移

链接后所有指令都有了全局唯⼀的地址。链接过程最重要的就是进⾏虚拟地址重定位(Relocation)

**编译**

`go tool compile -S ./hello.go | grep "hello.go:5|hello.go:6"` 

编译，即将源代码编译成 `.o` 目标文件，并输出汇编代码。 

**反编译**   https://golang.org/ref/spec go内部函数的用法

`go build main.go && go tool objdump ./main`

反汇编，即从可执行文件反编译成汇编，所以要先用 `go build` 命令编译出可执行文件。

**找到 runtime 源码**

例如，我想知道 go 关键字对应 runtime 里的哪个函数，于是写了一段测试代码：

```go
package main
 
func main() {
 go func() {
  println(1+2)
 }()
}
```

因为 go func(){}() 那一行代码在第 4 行，所以，grep 的时候加一个条件：

```cmd
go tool compile -S main.go | grep "main.go:4"
// 或
go build main.go && go tool objdump ./main | grep "main.go:4"
```
马上就能看到 `go func(){}()` 对应CALL  `newproc()` 函数，这时再深入研究下 `newproc()` 函数就大概知道 goroutine 是如何被创建的。 

**函数调⽤规约-函数栈**

局部变量只要不逃逸，都在栈上分配空间，为什么 Go 可以⼀个函数多个返回值，因为参数和返回值都是caller提供空间的

![1636634376105](F:\markdown笔记\Go高级工程师实战\image\1636634376105.png)

**初识 ast 的威力**

 在计算机科学中，抽象语法树（Abstract Syntax Tree，AST），或简称语法树（Syntax tree），是源代码语法结构的一种抽象表示。它以树状的形式表现编程语言的语法结构，树上的每个节点都表示源代码中的一种结构。 可以用 ast 包和 parser 包解析一个二元表达式 