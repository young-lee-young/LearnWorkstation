### Go并发问题

* 问题：并发赋值有并发安全问题

```go
package main

import (
	"fmt"
	"time"
)

func add(p *int32) {
	*p++
}

func main() {
	a := int32(0)

	for i := 0; i < 1000; i++ {
		go add(&a)
	}

	time.Sleep(time.Second * 1)

	fmt.Println(a)
}
```


* 解决：使用 atomic 解决

```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// sync/atomic/doc.go/AddInt32，汇编语言实现
func add(p *int32) {
	atomic.AddInt32(p, 1)
}

func main() {
	a := int32(0)

	for i := 0; i < 1000; i++ {
		go add(&a)
	}

	time.Sleep(time.Second * 1)

	fmt.Println(a)
}
```


* atomic 操作

1. 是一种硬件层面加锁机制
2. 保证操作一个变量的时候，其他协程、线程无法访问
3. 只能用于简单变量的简单操作


### CAS：compare and swap

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	c := int32(10)
	atomic.CompareAndSwapInt32(&c, 10, 100)
	fmt.Println(c)
}
```
