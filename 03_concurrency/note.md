# 并发编程
## 0.介绍

### 进程和线程
    进程是程序在操作系统中的一次执行过程，系统进行调度和资源分配的一个独立单位。
    线程是进程的一个执行实体，是CPU调度和分配的基本单位，他是比进程更小的能独立运行的基本单位。
    一个进程可以创建和撤销多个线程同一个进程中的多个线程可以并发执行。
### 并发和并行
    多线程程序在一个核的CPU上运行，是并发。
    多线程程序在多个核的CPU上运行，是并行。
### 协程和线程
    协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户线程的调度也是自己实现的。
    线程：一个线程上可以跑多个协程，协程是轻量级的线程。
### MIND
goroutine只是由官方实现的超级“线程池”。每个实例 4~5KB 的栈内存占用和由于实现机制而大幅减少的创造和销毁开销是go高并发的根本原因。

并发不是并行并发是在多个任务中进行切换，为什么要切换？是因为由于计算机CPU的计算速度非常快所以为了CPU的超强性能不被浪费而切换到下一项需要CPU的任务中。

## 1.goroutine
goroutine奉行通过通信来共享内存，而不是共享内存来通信。

    传统：围着数据转（抢，上锁）
    goroutine：跟着数据走

在Java/C++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量心智。

程序员需要一种只需要定义任务然后让系统分配任务到CPU上实现并发执行。

goroutine就是这样一种机制，goroutine类似于线程，但goroutine是由Go运行时产生的 **runtime** 调度和管理的。Go程序会智能地将goroutine中的任务合理地分配给每个CPU。GO之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

在Go语言中当需要让某个任务并发执行的时候，你只需要把任务包装成一个函数，开启一个goroutine去执行这个函数就行。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	fmt.Println("西瓜9毛一斤")
	time.Sleep(time.Second)
}
func hello() {
	fmt.Println("这瓜保甜吗")
}
```
如果将time.Sleep(time.Second*)去掉，程序会直接退出，不输出 这瓜保甜吗 因为在程序启动时，go程序就会为main()函数创建一个默认的goroutine。当mian()函数返回的时候goroutine就结束了，所有在main()函数中启动的goroutine也会一同结束，所以我么用time.Sleep来等一下hello()

还要注意的是 “西瓜9毛一斤” 在 “这瓜保甜吗”之前输出 因为创建goroutine 需要时间

### 启动多个goroutine
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	result := make(chan int, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(v int) { // 这里的 v 只是一个局部拷贝
			defer wg.Done()
			result <- v
		}(i) // 立即把当前的 i 传给参数 v
	}
	wg.Wait()
	close(result)
	for v := range result {
		fmt.Println(v)
	}
}

//如果有 100 个任务，但要求“每 3 个一组”有序执行，该怎么写？
//有 3 个工位（Goroutine），每个工位处理一部分逻辑，数据像传送带一样通过 Channel 流向下一个工位。这就是著名的 Pipeline（流水线）模式。用代码展示一下

```
### 可增长的栈
OS线程(操作系统线程)一般都有固定的栈内存1MB，一个goroutien栈内存在其生命周期开始时极小一般(2KB),goroutine 的栈内存极其不固定它可以按需增大和缩小，goroutine栈内存可达1GB
### GMP模型
GMP是go语言运行时(runtime)层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
- G(goroutine) 写下的 go func() 包含栈内存，指令指针等信息 初始约为 2KB
- M(Machine) 操作系统的物理线程 例如八核十六线程 每一个核映射两个线程 Goroutine最终要在Machine上运行
- P(Processor) 管理一组goroutine队列，P里会存储当前goroutine运行的上下文环境(函数指针，堆栈地址及地址边界) P会对自己管理的goroutine做一些调度(比如把占用CPU时间较长的goroutine暂停运行后续的goroutine) 当自己队列消费完了就去全局队列里取 如果全局队列里也消费完了回去其他的P队列里抢任务。

P与M一般也是一一对应的。P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，G所属的P会把G挂载在新建的M上 当旧的G阻塞完成或者认为其已经死掉时回收旧的M。

P的个数是通过runtime.GOMAXPROCS设定(最大为256个)，Go1.5版本后默认为物理线程数。在并发量大的时候会增加一些P和M，但不会太多，切换的太频繁的话会得不偿失。

但从线程上来讲，Goroutine是由runtime得调度器调度的，使用成为m:n调度的技术(复用/调度m个goroutine到n个OS线程)。其一大特点是goroutine的调度是在用户态下完成的，不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池，不直接调用系统对malloc函数(除非内存池也需要改变)，成本比调度OS线程低很多。另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，再加上goroutine的超轻量，以上种种保证了go调度的性能。

## 2.runtime包
在学runtime之前，常常说go程序是从func main()开始的，其实在运行程序时runtime会先接管CPU 初始化系统检查你的CPU核心数线程数据此设置GOMAXPROCS 然后启动调度器创建第一个线程M0和第一个协程GO 然后启动垃圾回收器(GC)

### runtime.Gosched()
让出CPU时间片
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("hello world") //主协程
	for i := 0; i < 2; i++ {
		//runtime.Gosched()
		fmt.Println("Hello")
	}
}
```
### runtime.Goexit()
退出当前协程
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			defer fmt.Println("C.defer")
			runtime.Goexit() //结束协程
			defer fmt.Println("D.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")

	}()
	time.Sleep(time.Second)//等待让goroutine能够创建成功
}

```
### runtime.GOMAXPROCS
Go与形式的调度器使用GOMAXPROCS参数来确定需要使用多少OS线程来同时执行Go代码。默认为CPU的核心数。
```go
//通过将任务分配到不同的CPU逻辑核心上实现并行效果
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	//runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second * 1)
}
```
goroutine和OS线程是多对多的关系，即m:n

## 3.Channel
单纯地将函数并发执行是没有意义的。函数与函数之间需要交换数据才能体现并发现执行函数的意义。

虽然可以使用共享内存的方式进行数据交换，但是共享内存在不同的goroutine中容易发生竞争问题。为了保证数据交换的正确性，必须使用互斥量堆内存进行加锁，这样的做法就会导致性能问题。

Go语言的并发模型是GSP(Communication Sequential Processes)，提倡 通过通信共享内存而不是通过共享内存而实现通信。

如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go语言中的channel是一种特殊的类型。通道像一个传送带或者队列，总是遵循先进先出的规则，保证发送数据的顺序。每一个通道都是一个具体类型的导管，也就是申明channel的时候需要为其指定元素类型。

 ```go
    //创建管道
    ch := make(chan int)    // 无缓冲管道
    ch := make(chan int, 5) // 有缓冲管道（容量为5）
	
	//如果是无缓冲管道下面的代码能通过编译但会在运行时报错 管道被阻塞
	//无缓冲管道也叫同步管道
	ch <- 20//发送数据
	a:=<-ch//接受数据
	close(ch)//关闭管道
```
当通过管道发送有限的数据时，我们可以通过close函数关闭通道来告知从该管道就收值的goroutine停止等待。当通道被关闭时，从该通道发送值会引发panic，从该通道里接受的值一直是类型零值。如何判断一个管道是否被关闭了呢？
```go
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for true {
			i, ok := <-ch1//用断言判断channel是否关闭
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	go func() {
		for i := range ch2 {
			fmt.Println(i) //通道关闭后自然退出
		}
	}()
}
```
### 单向管道
有时候我们会将channel作为参数在多个任务函数之间传递，很多时候我们在不同的任务函数中使用channel对其进行限制，比如限制channel在函数中只能发送或接受。
```go
package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
```


## 4.Goroutine池

## 5.定时器

## 6.select

## 7.并发安全和锁

## 8.