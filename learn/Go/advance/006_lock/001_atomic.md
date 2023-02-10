### 原子操作

* 问题：并发赋值有并发安全问题

```go
package main

import (
	"fmt"
	"time"
)

func add(p *int32) {
    // 并发有问题
	*p++
	
	// 解决：使用 atomic 解决并发问题
    // atomic.AddInt32(p, 1)
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
3. 只能用于**简单变量**的简单操作（所以需要实现更高级别的 mutex 锁）


* atomic.AddInt32（sync/atomic/doc.go/AddInt32，汇编语言实现）

```nasm
// runtime/internal/atomic/atomic_amd64.s/Xaddint32

TEXT ·Xadd(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), BX
	MOVL	delta+8(FP), AX
	MOVL	AX, CX
	LOCK                    // CPU 硬件锁
	XADDL	AX, 0(BX)
	ADDL	CX, AX
	MOVL	AX, ret+16(FP)
	RET

TEXT ·Xaddint32(SB), NOSPLIT, $0-20
	JMP	·Xadd(SB)
```


### CAS：Compare and Swap

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	c := int32(10)
	
	atomic.CompareAndSwapInt32(&c, 10, 100)
	// CompareAndSwapInt32 相当于
	// if c == 10 {
	//     c = 100
	// }
	
	fmt.Println(c)
}
```
