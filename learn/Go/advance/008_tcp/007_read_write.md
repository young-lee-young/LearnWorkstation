### read

* Read 方法

1. 直接调用 Socket 原生读方法
2. 如果失败，休眠等待可读
3. 被唤醒后再次调用系统 Socket

```go
// net/net.go/Read
package net

type conn struct {
	fd *netFD
}

// Implementation of the Conn interface.

// Read implements the Conn Read method.
func (c *conn) Read(b []byte) (int, error) {
	n, err := c.fd.Read(b)
	return n, err
}
```

```go
// net/fd_unix.go/Read
package net

func (fd *netFD) Read(p []byte) (n int, err error) {
	n, err = fd.pfd.Read(p)
}
```

```go
// internal/poll/fd_unix.go/
package poll

// Read implements io.Reader.
func (fd *FD) Read(p []byte) (int, error) {
    for {
    	// 系统调用读取数据
        n, err := syscall.Read(fd.Sysfd, p)
        if err != nil {
        	// 休眠等待数据
			if err == syscall.EAGAIN && fd.pd.pollable() {
				if err = fd.pd.waitRead(fd.isFile); err == nil {
					continue
				}
			}
        }
        err = fd.eofError(n, err)
        return n, err
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


### write

* Write

逻辑和 Read 完全相同
