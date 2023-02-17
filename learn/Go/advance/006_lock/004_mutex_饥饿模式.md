### 锁饥饿

* 什么是锁饥饿

协程一直尝试获取锁，始终获取不到


* 饥饿模式意义

1. 减少自旋竞争
2. 保证协程公平


* 饥饿模式工作原理

饥饿模式中，不自旋，新来的协程直接 sema 休眠

饥饿模式中，被唤醒的协程直接获取锁


* 进入饥饿模式

当前协程等待锁的时间超过 1 ms，切换到饥饿模式

```go
// runtime/mutex.go/lockSlow
package sync

const (
    starvationThresholdNs = 1e6
)

func (m *Mutex) lockSlow() {
	for {
	    // 如果被锁/饥饿模式，将等待协程数量加 1
	    // 饥饿模式新来的协程直接休眠，不进行自旋尝试
		if old&(mutexLocked|mutexStarving) != 0 {
            new += 1 << mutexWaiterShift
        }
		
		// The current goroutine switches mutex to starvation mode.
        // But if the mutex is currently unlocked, don't do the switch.
        // Unlock expects that starving mutex has waiters, which will not
        // be true in this case.
        // 将新状态设置成饥饿模式
        if starving && old&mutexLocked != 0 {
            new |= mutexStarving
        }
		
        if atomic.CompareAndSwapInt32(&m.state, old, new) {
	        // 协程没有获取到锁后，进入 sema 队列休眠
            runtime_SemacquireMutex(&m.sema, queueLifo, 1)
    
            // 其他协程释放锁后，协程被唤醒，计算休眠时间
            starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
        }
    }
}
```


* 退出饥饿模式

没有协程在 sema 队列中等待时，回到正常模式

```go
// runtime/mutex.go/lockSlow
package sync

func (m *Mutex) lockSlow() {
	for {
		// 饥饿模式
		if old&mutexStarving != 0 {
            delta := int32(mutexLocked - 1<<mutexWaiterShift)
            if !starving || old>>mutexWaiterShift == 1 {
                // Exit starvation mode.
                // Critical to do it here and consider wait time.
                // Starvation mode is so inefficient, that two goroutines
                // can go lock-step infinitely once they switch mutex
                // to starvation mode.
                delta -= mutexStarving
            }
            atomic.AddInt32(&m.state, delta)
            break
        }
	}
}
```


### 互斥锁使用经验

1. 减少锁使用时间
2. 善用 defer 确保锁的释放
