
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

	time.Sleep(time.Second)
	fmt.Println(a)
}
```

* 解决

使用 atomic 解决

```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func add(p *int32) {
	atomic.AddInt32(p, 1)
}

func main() {
	a := int32(0)

	for i := 0; i < 1000; i++ {
		go add(&a)
	}

	time.Sleep(time.Second)
	fmt.Println(a)
}
```


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
