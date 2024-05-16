### 阻塞模型和多路复用结合

```go
package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var body [100]byte

	for true {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("receive: %s\n", string(body[:]))

		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}
```
