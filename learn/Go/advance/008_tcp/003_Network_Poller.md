### pollDesc

* pollCache

```go
// runtime/netpoll.go/pollCache
package runtime

// 实际是 pollDesc 的链表头
type pollCache struct {
	lock  mutex     // 用来锁整个链表
	first *pollDesc
}
```


* pollDesc 结构体

pollDesc 指针是 Socket 相关详细信息，pollDesc 中记录了哪个协程休眠在等待此 Socket

```go
// runtime/netpoll.go/pollDesc 
package runtime

const (
	pdReady uintptr = 1
	pdWait  uintptr = 2
)

// 新增监听 Socket
// Network poller descriptor.
// 描述 Socket，并记录 Socket 和协程对应关系
type pollDesc struct {
	link *pollDesc // in pollcache, protected by pollcache.lock，指向下一个 pollDesc

	lock    mutex // protects the following fields
	fd      uintptr // Socket 文件描述
	
	rg      uintptr // pdReady, pdWait, G waiting for read or nil
	wg      uintptr // pdReady, pdWait, G waiting for write or nil
}
```


### NetWork Poller 初始化

1. poll_runtime_pollServerInit
2. 调用 netpollinit

```go
// runtime/netpoll.go/poll_runtime_pollServerInit
package runtime

func poll_runtime_pollServerInit() {
	netpollGenericInit()
}

func netpollGenericInit() {
	// 使用原子操作，保证只初始化一次
	// 一个 Go 程序只初始化一个 Network Poller
	lock(&netpollInitLock)

	// 初始化 epoll
	netpollinit()

	unlock(&netpollInitLock)
}
```


### Network Poller 新增监听 Socket

1. poll_runtime_pollOpen
2. 在 pollCache 链表中分配一个新的 pollDesc
3. 初始化 pollDesc（rg、wg 为 0）
4. 调用 netpollopen

```go
// runtime/netpoll.go/poll_runtime_pollOpen
package runtime

// fd：Socket 的文件描述符
func poll_runtime_pollOpen(fd uintptr) (*pollDesc, int) {
	// 分配新的 pollDesc
	pd := pollcache.alloc()

	// 初始化 pollDesc
	lock(&pd.lock) pd.rg = 0
	pd.wg = 0
    unlock(&pd.lock)

	// 向 epoll 新增一个事件
	errno := netpollopen(fd, pd)

	return pd, 0
}

func (c *pollCache) alloc() *pollDesc { 

}
```
