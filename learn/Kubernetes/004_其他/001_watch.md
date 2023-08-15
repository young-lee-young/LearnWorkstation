* 服务端 watch 示例

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	addr     = ":8000"
	notifyCh = make(chan int)
)

func production() {
	for i := 0; i < 5; i++ {
		notifyCh <- i + 1
		<-time.Tick(time.Second * 3)
	}
	close(notifyCh)
}

func watch(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)
	for {
		v, ok := <-notifyCh
		if !ok {
			return
		}
		// 写入到发送缓存
		fmt.Fprintf(w, "%v\n", v)
        // 这里会真正的发送到客户端
		flusher.Flush()
	}
}

func main() {
	go production()

	http.HandleFunc("/watch", watch)
	
	_ = http.ListenAndServe(addr, nil)
}
```


* 客户端请求

服务端返回的 Transfer-Encoding: chunked 会使客户端不断开连接，等待服务端持续推送数据

```bash
$ curl http://127.0.0.1:8000/watch -vvv
*   Trying 127.0.0.1:8000...
* Connected to 127.0.0.1 (127.0.0.1) port 8000 (#0)
> GET /watch HTTP/1.1
> Host: 127.0.0.1:8000
> User-Agent: curl/7.87.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 29 Jun 2023 02:16:41 GMT
< Content-Type: text/plain; charset=utf-8
< Transfer-Encoding: chunked
<
1
2
3
4
5
* Connection #0 to host 127.0.0.1 left intact
```
