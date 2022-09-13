# Go语言运行时


### runtime基础

* 什么是runtime

为了实现额外的功能，而在程序运行时自动加载/运行的一些模块


* runtime包括以下模块

scheduler：管理所有的G、M、P，在后台执行调度循环

netpoll：网络轮询，负责管理网络FD相关读写

memory：当代码需要内存时，负责内存分配工作

garbage：当内存不再需要时，负责回收内存


* runtime代码位置

/usr/local/go/src/runtime


### G、M、P模型

G：表示goroutine，每一次go func() {}都是创建了一个G

M：表示线程，消费G

P：维护存放G的一些本地队列，个数为GOMAXPROCS个数

* 队列优先级

多级队列较少锁竞争

runnext > local run queue（数组） > global run queue（链表）

* G生产过程

创建G -> 进入runnext -> 尝试将在runnext中的旧的G放在本地队列 -> 本地队列未满，放在本地队列 -> 如果本地队列已满(256)，将本地队列拿出一半 + 旧runnext G，放在全局队列

* G消费过程

1. 每60次去全局队列取一个G，防止全局队列G饿死
2. runnext和本地队列
3. 从全局队列拉一半到本地队列，加全局锁
4. 从其他goroutine中偷一半到本地队列，无锁


### 非阻塞情况

1. 管道

```go
// 向没有buffer的管道写入数据
var ch = make(chan int)
ch <- 1

// 从没有数据的管道收数据
var ch = make(chan int)
<- ch
``` 

2. 睡眠

```go
time.Sleep(time.Second * 1)
```

3. 网络读写

```go
// 网络读
var c net.Conn
var buf = make([]byte, 1024)

// data not ready, block here
n, err := c.Read(buf)

// 网络写
var c net.Conn
var buf = []byte("hello")

// send buffer full, white blocked
n, err := c.Write(buf)
```

4. select语句

```go
var (
    ch1 = make(chan int)
    ch2 = make(chan int)
)

// no case ready, block
select {
case <- ch1:
    fmt.Println("ch1 ready")
case <- ch2:
    fmt.Println("ch2 ready")
}
```

5. 锁操作

```go
var m sync.RWMutex

// somebody already grab the lock, block here
m.lock()
```

以上这些情况不会阻塞调度循环，而是把goroutine挂起

所谓的挂起，就是让goroutine先进某个数据结构，待ready后再继续执行，不会占用线程

此时，线程会进入schedule，继续消费队列，执行其他goroutine


### 阻塞情况

1. 执行CGo

在执行C代码，或者阻塞在syscall上时，必须占用一个线程
