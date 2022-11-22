### Network Poller 收发数据

分为两个场景：

1. 协程需要收发数据时，Socket 已经可读写
2. 协程需要收发数据时，Socket 暂时无法读写


### GC 调用 netpoll

```go
// runtime/mgc.go/gcStart
package runtime

func gcStart(trigger gcTrigger) {
    systemstack(func() {
        now = startTheWorldWithSema(trace.enabled)
    })
}
```

```go
// runtime/proc.go/startTheWorldWithSema
package runtime

func startTheWorldWithSema(emitTraceEvent bool) int64 {
	if netpollinited() {
		// 调用 netpoll
		list := netpoll(false)
	}
}
```


### Socket 已经可读写

* Go 将 pollDesc rg/wg 置为 pdReady

runtime 的 GC 循环调用 netpoll() 方法（g0协程），netpoll 系统调用 poll_wait，发现 Socket 可读写时，给对应的 rg/wg 置为 pdReady（1）

```go
// runtime/netpoll_epoll.go/netpoll
package runtime

const (
	pdReady uintptr = 1
)

func netpoll(delay int64) gList {
	// 可以读写
    if mode != 0 {
        pd := *(**pollDesc)(unsafe.Pointer(&ev.data))
        // 设置为 pdReady
        netpollready(&toRun, pd, mode)
    }
}
```

```go
// runtime/netpoll.go/netpollready
package runtime

func netpollready(toRun *gList, pd *pollDesc, mode int32) {
    var rg, wg *g
    if mode == 'r' || mode == 'r'+'w' {
        rg = netpollunblock(pd, 'r', true)
    }
    if mode == 'w' || mode == 'r'+'w' {
        wg = netpollunblock(pd, 'w', true)
    }
}

func netpollunblock(pd *pollDesc, mode int32, ioready bool) *g {
	// 获取当前的 rg/wg
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}

	for {
		old := gpp.Load()

		var new uintptr
		// Socket 可以读写，将 rg/wg 置为 pdReady（1）
		if ioready {
			new = pdReady
		}

		// 设置为 pdReady
		if gpp.CompareAndSwap(old, new) {

		}
	}
}
```


* 业务协程调用 poll_wait，判断 pollDesc rg/wg 状态

业务协程想要读写 Socket 时，调用 poll_runtime_pollWait()，判断 rg/wg 已经置为 pdReady（1），返回 0

```go
// runtime/netpoll.go/poll_runtime_pollWait
package runtime

const (
	pollNoError = 0 // no error
)

// runtime/netpoll.go/poll_runtime_pollWait()
func poll_runtime_pollWait(pd *pollDesc, mode int) int {
	// 当 pdReady 时，返回 ture，取非后，退出 for 循环
	for !netpollblock(pd, int32(mode), false) {

	}

	// 返回 0，no error
	return pollNoError
}

func netpollblock(pd *pollDesc, mode int32, waitio bool) bool {
	// 获取当前的 rd/wg
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}

	for {
		// 判断当前 rg/wg 是否为 pdReady
		if gpp.CompareAndSwap(pdReady, 0) {
			return true
		}
	}
}
```


### Socket 暂时无法读写

* Go 不修改 pollDesc rg/wg

runtime 的 GC 循环调用 netpoll() 方法（g0协程），发现没有 ready 的 Socket，不修改任何值，一直重试

```go
// runtime/netpoll_epoll.go/netpoll
package runtime

func netpoll(delay int64) gList {
	// epoll_wait系统调用，没有拿到 ready 的 Socket，一直重试
    n := epollwait(epfd, &events[0], int32(len(events)), waitms)
}
```

如果发现 Socket 可读写，将可读写协程列表返回给 runtime，由 runtime 唤起协程，进行数据处理


* 业务方法调用 epoll_wait，没有 ready 的 Socket，休眠等待

业务协程调用 poll_runtime_pollWait()，发现 rg/wg 为0，给对应的 rg/wg 置为协程地址，休眠等待

```go
// runtime/netpoll.go/poll_runtime_pollWait
package runtime

const (
	pdWait uintptr = 2 // 有人在操作 rg/wg
)

func poll_runtime_pollWait(pd *pollDesc, mode int) int {
	for !netpollblock(pd, int32(mode), false) {

	}
}

func netpollblock(pd *pollDesc, mode int32, waitio bool) bool {
	// 获取当前的 rd/wg
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}

	for {
		// 将 rg/wg 改为 pdWait，跳出循环
		if gpp.CompareAndSwap(0, pdWait) {
			break
		}
	}

	// 调用 gopark 休眠协程
	// gpp 此时为 2，在 gopark 中，会调用 netpollblockcommit，将当前协程的地址赋值给 gpp
    if waitio || netpollcheckerr(pd, mode) == 0 {
	    gopark(netpollblockcommit, unsafe.Pointer(gpp), waitReasonIOWait, traceEvGoBlockNet, 5)
    }
}

func netpollblockcommit(gp *g, gpp unsafe.Pointer) bool {
	// 将当前协程的地址赋值到 gpp 上
	r := atomic.Casuintptr((*uintptr)(gpp), pdWait, uintptr(unsafe.Pointer(gp)))
	return r
}
```
