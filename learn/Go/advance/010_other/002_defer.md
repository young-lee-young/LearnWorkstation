### defer 实现

1. 协程记录 defer 信息，函数退出时调用
2. 将 defer 代码直接编译进函数尾


### 协程记录 defer 信息，函数退出时调用

* 堆上分配（1.12 之前）

在堆上开辟一个 sched.deferpool，遇到 defer 语句，将信息放入 deferpool，函数返回时，从 deferpool 取出执行

```go
package runtime

// 协程结构体（runtime/runtime2.go/g）
type g struct {
	_defer       *_defer
}

// 协程队列结构体（runtime/runtime2.go/p）
type p struct {
	// 记录一些 defer
	deferpool    []*_defer // pool of available defer structs (see panic.go)
	deferpoolbuf [32]*_defer
}

// defer 结构体（runtime/runtime2.go/_defer）
type _defer struct {
	started bool
	heap    bool

	openDefer bool
	sp        uintptr
	pc        uintptr
	fn        func()
	_panic    *_panic
	link      *_defer

	fd   unsafe.Pointer
	varp uintptr

	framepc uintptr
}
```

```go
package runtime

// 函数返回是调用 deferreturn
func deferreturn() {
	// 获取当前的 g
	gp := getg()
	for {
		// 获取 defer
		d := gp._defer
		// 调用 defer 方法
		done := runOpenDeferFrame(gp, d)
        // 释放 defer
		freedefer(d)
	}
}
```

* 栈上分配（1.13 之后）

遇到 defer 语句，将信息放入栈上，函数返回时，从栈中取出执行

但是栈上分配只能保存一个 defer 信息


### 将 defer 代码直接编译进函数尾

* 开放编码（1.14 之后）

如果 defer 语句在编译时可以固定，直接改用户代码，defer 语句放入函数末尾
