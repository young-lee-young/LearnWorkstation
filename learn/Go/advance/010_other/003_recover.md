# panic 和 recover


### panic

panic 会抛出错误，终止协程运行，带崩整个 Go 程序

* 示例

```go
package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		panic("panic ----------")
		// 下面的代码都不会执行
		fmt.Println("end g")
	}()

	time.Sleep(time.Second)

    // 执行不到
	fmt.Println("end main g")
}
```


### panic + defer

panic 在退出协程之前会执行所有已注册的 defer

但是不会执行其他协程的 defer

* 示例

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 执行不到
	defer fmt.Println("defer main g")

	go func() {
		// 会执行到
		defer fmt.Println("defer g")
		panic("panic -------------")
		// 下面的代码都不会执行
		fmt.Println("end g")
	}()

	time.Sleep(time.Second)
	
	fmt.Println("end main g")
}
```


* 源码

```go
// runtime/panic.go/gopanic
package runtime

func gopanic(e any) {
	gp := getg()

	for {
		// 获取当前协程的 defer，进行处理
		// 但是不会获取其他携程的 defer
		d := gp._defer
	}
}
```


### panic + defer + recover

* 原理

1. 如果涉及 recover，defer 会使用堆上分配（deferpool）
2. 遇到 panic，panic 会从 deferpool 取出 defer 语句执行
3. defer 中调用 recover，可以终止 panic 过程（上一级的协程不会崩溃）


* 示例

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 会执行到
	defer fmt.Println("defer main g")

	go func() {
		defer func() {
			recover()
		}()
		panic("panic ----------------")
		// 下面的代码都不会执行
		fmt.Println("end g")
	}()

	time.Sleep(time.Second)
	
	// 会执行到
	fmt.Println("end main g")
}
```

* 源码

```go
// runtime/panic.go/gopanic
package runtime

func gopanic(e any) {
	gp := getg()

	for {
		d := gp._defer
		
		// 执行 recover 操作，恢复协程
		if p.recovered {
			
		}
	}
}
```
