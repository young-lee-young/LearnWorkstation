### 锁饥饿

* 什么是锁饥饿

协程一直尝试获取锁，始终获取不到


* 进入饥饿模式

当前协程等待锁的时间超过 10ms，切换到饥饿模式


* 退出饥饿模式

没有协程在 sema 队列中等待时，回到正常模式


* 饥饿模式工作原理

饥饿模式中，不自旋，新来的协程直接 sema 休眠

饥饿模式中，被唤醒的协程直接获取锁

```go
// runtime/mutex.go/lockSlow
package sync

func (m *Mutex) lockSlow() {
	for {
		// 退出饥饿模式
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
