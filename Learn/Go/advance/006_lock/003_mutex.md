# Mutex 互斥锁


### 结构体

```go
package sync

type Mutex struct {
	state int32
	sema  uint32
}
```


### 饥饿模式

当前协程等待锁的时间超过 1ms，切换到饥饿模式

饥饿模式中，不自旋，新来的协程直接 sema 休眠

饥饿模式中，被唤醒的协程直接获取锁

没有协程在队列中等待时，回到正常模式


### 锁使用经验

1. 减少锁使用时间
2. 善用 defer 确保锁的释放
