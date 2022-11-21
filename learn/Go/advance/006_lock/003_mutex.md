# Mutex 互斥锁


### 并发问题

* 问题

```go
package main

import (
	"fmt"
	"time"
)

type Person struct {
	salary int
	level  int
}

func (p *Person) promote() {
	p.salary += 1
	fmt.Println(p.salary)
	p.level += 1
	fmt.Println(p.level)
}

func main() {
    a := Person{salary: 0, level: 0}

	for i := 0; i < 1000; i++ {
		go a.promote()
	}

	time.Sleep(time.Second * 10)
	fmt.Println(a.level, a.salary)
}
```


* 解决

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu     sync.Mutex   // 加互斥锁
	salary int
	level  int
}

func (p *Person) promote() {
	p.mu.Lock()
	p.salary += 1
	fmt.Println(p.salary)
	p.level += 1
	fmt.Println(p.level)
	p.mu.Unlock()
}

func main() {
	a := Person{salary: 0, level: 0}

	for i := 0; i < 1000; i++ {
		go a.promote()
	}

	time.Sleep(time.Second * 10)
	fmt.Println(a.level, a.salary)
}
```


### 结构体

![Mutex 锁结构](images/mutex锁结构.png)

```go
// sync/mutex.go/Mutex
package sync

// A Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.
type Mutex struct {
	state int32
	sema  uint32
}
```

state：共 32 位
1 - 29：WaiterShift，等待锁的协程数量
30：Starving，饥饿模式
31：Woken，唤醒
32：Locked，0 没有被锁，1 被锁住


### 正常模式加锁

1. 尝试 CAS 直接加锁
2. 若无法直接获取，进行多次**自旋**尝试
3. 多次尝试失败，进入 sema 队列休眠

```go
// runtime/mutex.go/Lock
package sync

const (
	mutexLocked = 1 << iota // mutex is locked
)

func (m *Mutex) Lock() {
	// 尝试加锁：将 state 从 0 置为 1
    if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
	
    }
    
    // 加锁失败
    m.lockSlow()
}

func (m *Mutex) lockSlow() {
	for {
		// 判断是否是上锁状态或饥饿状态 && 可以自旋
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// 执行自旋
			runtime_doSpin()
			continue
		}
		
		// 协程进入 sema 队列休眠
        runtime_SemacquireMutex(&m.sema, queueLifo, 1)
		
		// 协程被唤醒后，会接着执行，再回到 for 循环里
        awoke = true
	}
}
```


### 正常模式解锁

1. 尝试 CAS 直接解锁
2. 若发现有协程在 sema 中休眠，唤醒一个协程

```go
// runtime/mutex.go/Unlock
package sync

const (
	mutexLocked = 1 << iota // mutex is locked
)

func (m *Mutex) Unlock() {
	// 解锁：将状态置为 0
   	new := atomic.AddInt32(&m.state, -mutexLocked)
   	// 等待的协程数量是否为 0，不为 0，唤醒协程
   	if new != 0 {
		m.unlockSlow(new)
	}
}

func (m *Mutex) unlockSlow(new int32) {
	// Grab the right to wake someone.
	new = (old - 1<<mutexWaiterShift) | mutexWoken
	if atomic.CompareAndSwapInt32(&m.state, old, new) {
		// 将协程从 sema 队列释放
		runtime_Semrelease(&m.sema, false, 1)
		return
	}
}
```
