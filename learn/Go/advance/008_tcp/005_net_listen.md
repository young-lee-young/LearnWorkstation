### net 包

net 包是 go 原生的网络包，实现了 TCP、UDP、HTTP 等网络操作

* Listener 接口

```go
// net/net.go/Listener
package net

type Listener interface {
	// Accept waits for and returns the next connection to the listener.
	Accept() (Conn, error)

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() Addr
}
```


* Listen 方法

1. 新建 Socket，并执行 bind 操作
2. 新建一个 FD（net 包对 Socket 的详细描述）
3. 将 TCPListener 的 FD 信息加入监听
4. 返回一个 TCPListener 对象（TCPListener 本质上是一个 LISTEN 状态的 Socket）

```go
// net/dial.go/Listen
package net

func Listen(network, address string) (Listener, error) {
	var lc ListenConfig
	return lc.Listen(context.Background(), network, address)
}

func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error) {
    l, err = sl.listenTCP(ctx, la)
}
```

```go
// net/tcpsock_posix.go/listenTCP
package net

// TCPListener is a TCP network listener. Clients should typically
// use variables of type Listener instead of assuming TCP.
type TCPListener struct {
	fd *netFD
	lc ListenConfig
}

func (sl *sysListener) listenTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
	fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_STREAM, 0, "listen", sl.ListenConfig.Control)
	if err != nil {
		return nil, err
	}
	return &TCPListener{fd: fd, lc: sl.ListenConfig}, nil
}
```

```go
// net/ipsock_posix.go/internetSocket
package net

func internetSocket(ctx context.Context, net string, laddr, raddr sockaddr, sotype, proto int, mode string, ctrlFn func(string, string, syscall.RawConn) error) (fd *netFD, err error) {
	return socket(ctx, net, family, sotype, proto, ipv6only, laddr, raddr, ctrlFn)
}
```

```go
// net/sock_posix.go/socket
package net

// socket returns a network file descriptor that is ready for
// asynchronous I/O using the network poller.
func socket(ctx context.Context, net string, family, sotype, proto int, ipv6only bool, laddr, raddr sockaddr, ctrlFn func(string, string, syscall.RawConn) error) (fd *netFD, err error) {
	// 系统调用创建 Socket
	s, err := sysSocket(family, sotype, proto)
	
	// 新建 FD
    if fd, err = newFD(s, family, sotype, net); err != nil {
        return nil, err
    }
    
    
    switch sotype {
    case syscall.SOCK_STREAM, syscall.SOCK_SEQPACKET:
        if err := fd.listenStream(laddr, listenerBacklog(), ctrlFn); err != nil {
            fd.Close()
            return nil, err
        }
        return fd, nil
    }
}

func (fd *netFD) listenStream(laddr sockaddr, backlog int, ctrlFn func(string, string, syscall.RawConn) error) error {
    if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
        return os.NewSyscallError("bind", err)
    }
    
    if err = fd.init(); err != nil {
        return err
    }
}
```

```go
// net/fd_unix.go/newFD
package net

// Network file descriptor.
type netFD struct {
	pfd poll.FD
}

func newFD(sysfd, family, sotype int, net string) (*netFD, error) {
	ret := &netFD{
		pfd: poll.FD{
			Sysfd:         sysfd,
			IsStream:      sotype == syscall.SOCK_STREAM,
			ZeroReadIsEOF: sotype != syscall.SOCK_DGRAM && sotype != syscall.SOCK_RAW,
		},
		family: family,
		sotype: sotype,
		net:    net,
	}
	return ret, nil
}

func (fd *netFD) init() error {
	return fd.pfd.Init(fd.net, true)
}
```


* FD 结构体

```go
// internal/poll/fd_unix.go/FD
package poll

// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
type FD struct {
	// System file descriptor. Immutable until Close.
	Sysfd int

	// I/O poller.
	pd pollDesc
}

func (fd *FD) Init(net string, pollable bool) error {
    err := fd.pd.init(fd)
}
```

```go
// internal/poll/fd_poll_runtime.go/init
package poll

var serverInit sync.Once

func (pd *pollDesc) init(fd *FD) error {
    serverInit.Do(runtime_pollServerInit)
    ctx, errno := runtime_pollOpen(uintptr(fd.Sysfd))
}
```
