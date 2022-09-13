### NetWork Poller

* pollCache（runtime/netpoll.go/pollCache）

```go
package runtime

// 实际是 pollDesc 的链表头
type pollCache struct {
	lock  mutex // 用来锁整个链表
	first *pollDesc
}
```

* pollDesc（runtime/netpoll.go/pollDesc）

```go
package runtime

const (
	pdReady uintptr = 1
	pdWait  uintptr = 2
)

// 描述 Socket，并记录 Socket 和协程对应关系
type pollDesc struct {
	// 指向下一个 pollDesc
	link *pollDesc // in pollcache, protected by pollcache.lock
	fd   uintptr   // Socket 的文件描述符

	atomicInfo atomic.Uint32

	rg atomic.Uintptr // pdReady, pdWait, G waiting for read or nil
	wg atomic.Uintptr // pdReady, pdWait, G waiting for write or nil

	lock    mutex // protects the following fields
	closing bool
	user    uint32    // user settable cookie
	rseq    uintptr   // protects from stale read timers
	rt      timer     // read deadline timer (set if rt.f != nil)
	rd      int64     // read deadline (a nanotime in the future, -1 when expired)
	wseq    uintptr   // protects from stale write timers
	wt      timer     // write deadline timer
	wd      int64     // write deadline (a nanotime in the future, -1 when expired)
	self    *pollDesc // storage for indirect interface. See (*pollDesc).makeArg.
}
```


### NetWork Poller 初始化（runtime/netpoll.go/poll_runtime_pollServerInit()）

```go
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


### Network Poller 新增监听 Socket（runtime/netpoll.go/poll_runtime_pollOpen()）

```go
package runtime

// 新增监听 Socket
// fd：Socket 的文件描述符
func poll_runtime_pollOpen(fd uintptr) (*pollDesc, int) {
	// 分配 pollDesc
	pd := pollcache.alloc()

	// 向 epoll 新增一个事件
	errno := netpollopen(fd, pd)

	return pd, 0
}
```
