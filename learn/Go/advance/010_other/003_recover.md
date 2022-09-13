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
	}()

	time.Sleep(time.Second)
}
```

```go
package runtime

// runtime/panic.go/gopanic()
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
	}()

	time.Sleep(time.Second)
	// 会执行到
	fmt.Println("end main g")
}
```

```go
package runtime

// runtime/panic.go/gopanic()
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
