### 系统指令集

cmd/vendor/golang.org/x/sys/unix/zsysnum_linux_amd64.go

```go
package unix

const (
	SYS_EPOLL_CREATE  = 213
	SYS_EPOLL_CREATE1 = 291
)
```


### epoll 多路复用

* 实现的功能

epoll_create()：新建多路复用

epoll_ctl()：往多路复用器里插入需要监听的事件

epoll_wait()：查询发生的事件


### Go epoll

* 抽象

epoll 抽象层是为了统一各个操作系统对多路复用器的实现

* netpollinit（runtime/netpoll_epoll.go/netpollinit()）

```go
package runtime

// 对于操作系统 epoll 的描述
var epfd int32 = -1 // epoll descriptor

// 初始化 epoll
func netpollinit() {
	// 创建 epoll
	epfd = epollcreate1(_EPOLL_CLOEXEC)
	if epfd < 0 {
		epfd = epollcreate(1024)
	}

	// 初始化一个管道，管道用来关闭 epoll
	r, w, errno := nonblockingPipe()

	// 往 epoll 里添加事件
	errno = epollctl(epfd, _EPOLL_CTL_ADD, r, &ev)
}
```

* epoll_create（runtime/netpoll_epoll.go/epollcreate()，汇编实现）

```go
package runtime

// 创建 epoll
func epollcreate(size int32) int32

func epollcreate1(flags int32) int32
```

实际实现（runtime/sys_linux_amd64.go/runtime.epollcreate）

```asm
// int32 runtime·epollcreate(int32 size);
TEXT runtime·epollcreate(SB),NOSPLIT,$0
    // 系统调用，创建 epoll
	MOVL    $SYS_epoll_create, AX
	SYSCALL
	RET

// int32 runtime·epollcreate1(int32 flags);
TEXT runtime·epollcreate1(SB),NOSPLIT,$0
    // 系统调用，创建 epoll
	MOVL	$SYS_epoll_create1, AX
	SYSCALL
	RET

```

* netpollopen（runtime/netpoll_epoll.go/netpollopen()）

```go
package runtime

// 增加一个事件
// 传入 Socket 的 fd，和 pollDesc 的指针
// pollDesc 中记录了哪个协程休眠等待此 Socket
func netpollopen(fd uintptr, pd *pollDesc) int32 {
	var ev epollevent
	// 四个事件
	ev.events = _EPOLLIN | _EPOLLOUT | _EPOLLRDHUP | _EPOLLET
	*(**pollDesc)(unsafe.Pointer(&ev.data)) = pd

	// 往 epoll 里添加事件
	return -epollctl(epfd, _EPOLL_CTL_ADD, int32(fd), &ev)
}
```

* netpoll（runtime/netpoll_epoll.go/netpoll()）

```go
package runtime

type gList struct {
	head guintptr
}

// netpoll checks for ready network connections
// 检查是否有事件发生
func netpoll(delay int64) gList {
	var events [128]epollevent

	// 系统调用 epollwait，初始化长度为 128 的数组用来接收事件
	n := epollwait(epfd, &events[0], int32(len(events)), waitms)

	// 没有事件发生
	if n < 0 {
		// 重试
	}

	// 协程列表
	var toRun gList
	for i := int32(0); i < n; i++ {
		// 读写标志位
		var mode int32
		if ev.events&(_EPOLLIN|_EPOLLRDHUP|_EPOLLHUP|_EPOLLERR) != 0 {
			mode += 'r'
		}
		if ev.events&(_EPOLLOUT|_EPOLLHUP|_EPOLLERR) != 0 {
			mode += 'w'
		}

		if mode != 0 {
			// 找到了 Socket 和 协程描述
			pd := *(**pollDesc)(unsafe.Pointer(&ev.data))

			// 将协程放入
			netpollready(&toRun, pd, mode)
		}
	}

	// 返回可读、可写协程列表
	return toRun
}

// 判断 Socket 可读、可写情况
func netpollready(toRun *gList, pd *pollDesc, mode int32) {
	var rg, wg *g
	if mode == 'r' || mode == 'r'+'w' {
		rg = netpollunblock(pd, 'r', true)
	}
	if mode == 'w' || mode == 'r'+'w' {
		wg = netpollunblock(pd, 'w', true)
	}

	// 放入可读列表
	if rg != nil {
		toRun.push(rg)
	}

	// 放入可写列表
	if wg != nil {
		toRun.push(wg)
	}
}
```
