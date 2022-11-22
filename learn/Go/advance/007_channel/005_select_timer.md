### select 原理

1. 同时存在接收、发送、默认路径
2. 首先查看是否有可以立即执行的 case
3. 没有的话，有 default，走 default
4. 没有 default，把自己（即执行 select 的协程）注册在**所有的** channel 中，休眠等待


### timer

可以提供一个 channel，定时塞入数据

```go
package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTimer(time.Second)

	i := <- t.C
	
	fmt.Println(i)
}
```
