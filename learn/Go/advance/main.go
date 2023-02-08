package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	time.Sleep(time.Second)

	close(ch)

	time.Sleep(time.Second)
}

func stopServer() {
	handler := new(HelloHandler)

	server := &http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: handler,
	}

	go graceful(server)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("main goroutine exit")
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 10)
	w.WriteHeader(200)
	w.Write([]byte("Hello"))
}

/**
优雅停止服务
*/
func graceful(server *http.Server) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	done := make(chan bool, 1)

	go func() {
		sig := <-c

		fmt.Println(sig)

		done <- true
	}()

	fmt.Println("waiting signal")
	<-done
	fmt.Println("exit")
}

/**
控制 goroutine 数量
*/
func goroutineCount() {
	var wg sync.WaitGroup

	ch := make(chan struct{}, 10)

	for i := 0; i < math.MaxInt64; i++ {
		wg.Add(1)

		ch <- struct{}{}

		go func(i int) {
			defer wg.Done()

			fmt.Println("i:", i)

			time.Sleep(time.Second * 1)

			<-ch
		}(i)
	}

	wg.Wait()
}

/**
高效字符串拼接
*/
func stringBuilder() {
	var sb strings.Builder

	for i := 0; i < 100; i++ {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("-")
	}

	ret := sb.String()

	fmt.Println("ret", ret)
}

/**
正则匹配

大于等于6，小于等于15长度，首字母是字符，由下划线、数字、字母组成
*/
func reg() {
	// 5 - 14
	pattern := "^[a-zA-Z][_0-9a-zA-Z]{5,14}"

	//text := "leeyoung"
	text := ""

	match, _ := regexp.MatchString(pattern, text)

	fmt.Println("match:", match)
}
