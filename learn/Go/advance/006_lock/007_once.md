### 举例

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	salary int
	level  int
}

func (p *Person) promote() {
	p.salary += 1
	p.level += 1
}

func main() {
	a := Person{salary: 0, level: 0}

	once := sync.Once{}
	go once.Do(a.promote)
	go once.Do(a.promote)
	go once.Do(a.promote)

	time.Sleep(time.Second * 1)
	
	fmt.Println(a.level, a.salary)
}
```

### 原理

1. 争抢一个 mutex，抢不到的陷入 sema 休眠
2. 抢到的执行代码，改值，释放锁
3. 其他协程唤醒后判断值已经修改，直接返回


### 结构

* 结构体

```go
// sync/once.go/Once
package sync

type Once struct {
	done uint32
	m    Mutex
}
```


### 源码

```go
package sync

func (o *Once) Do(f func()) {
	// 如果等于 0，方法还没做，可以做
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	// 加锁
    o.m.Lock()
	defer o.m.Unlock()
    
	if o.done == 0 {
		// 修改 done 的值
		defer atomic.StoreUint32(&o.done, 1)
		// 执行业务方法
		f()
	}
}
```
