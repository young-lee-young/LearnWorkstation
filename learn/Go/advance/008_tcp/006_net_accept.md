### Accept 

* 获取新连接的方法：两种方法

1. 直接调用系统 accept 检查是否有新的连接
2. Listen 的 Socket 已经加入到 Network Poller 里面去监听新连接，可以通过 Network Poller 检查是否有新连接


* Accept 方法

1. 直接调用 Socket 的 accept 方法
2. 如果失败，休眠等待新的连接
3. 将新的 Socket FD 加入监听，并且将 Socket FD 包装成 TCPConn 返回（TCPConn 本质上是一个 ESTABLISHED 状态的 Socket）

```go
// net/net.go/conn
package net

type conn struct {
	fd *netFD
}
```

```go
// net/tcpsock.go/Accept
package net

// TCPListener is a TCP network listener. Clients should typically
// use variables of type Listener instead of assuming TCP.
type TCPListener struct {
	fd *netFD
	lc ListenConfig
}

type TCPConn struct {
	conn // net 包里的 conn
}

// Accept implements the Accept method in the Listener interface; it
// waits for the next call and returns a generic Conn.
func (l *TCPListener) Accept() (Conn, error) {
	c, err := l.accept()
	return c, nil
}

func newTCPConn(fd *netFD) *TCPConn {
	c := &TCPConn{conn{fd}}
	return c
}
```

```go
// net/tcpsock_posix.go/accept
package net

func (ln *TCPListener) accept() (*TCPConn, error) {
	fd, err := ln.fd.accept()
	
	tc := newTCPConn(fd)

	return tc, nil
}
```

```go
// net/fd_unix.go/accept
package net

func (fd *netFD) accept() (netfd *netFD, err error) {
    // 获取连接的 Socket 的 FD
    d, rsa, errcall, err := fd.pfd.Accept()
    
    // 将系统的 FD 包装成 Go 的 netFD
    if netfd, err = newFD(d, fd.family, fd.sotype, fd.net); err != nil {
        return nil, err
    }
    
    // 将新的 Socket FD 加入到多路复用里
    if err = netfd.init(); err != nil {
        fd.Close()
        return nil, err
    }
    
    return netfd, nil
}
```

```go
// internal/poll/fd_unix.go/Accept
package poll

func (fd *FD) Accept() (int, syscall.Sockaddr, string, error) {
    for {
    	// 第一种方法：accept 系统调用，获取连接的 Socket
        s, rsa, errcall, err := accept(fd.Sysfd)
        if err == nil {
            return s, rsa, "", err
        }
        
        switch err {
        case syscall.EAGAIN:
            // 第二种方法：Network Poller 获取新连接
            if err = fd.pd.waitRead(fd.isFile); err == nil {
                continue
            }
        }
    }
}
```

```go
// internal/poll/fd_poll_runtime.go/wait
package poll

func (pd *pollDesc) wait(mode int, isFile bool) error {
	res := runtime_pollWait(pd.runtimeCtx, mode)
}

func (pd *pollDesc) waitRead(isFile bool) error {
	return pd.wait('r', isFile)
}
```
