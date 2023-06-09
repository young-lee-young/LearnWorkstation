//////package main
//////
//////import (
//////	"fmt"
//////	"log"
//////	"math"
//////	"net/http"
//////	"os"
//////	"os/signal"
//////	"regexp"
//////	"strconv"
//////	"strings"
//////	"sync"
//////	"syscall"
//////	"time"
//////)
//////
//////
//////func main() {
//////	var s []int
//////	s[0] = 0
//////
//////	s2 := make([]int, 0)
//////}
//////
//////func stopServer() {
//////	handler := new(HelloHandler)
//////
//////	server := &http.Server{
//////		Addr:    "127.0.0.1:8001",
//////		Handler: handler,
//////	}
//////
//////	go graceful(server)
//////
//////	err := server.ListenAndServe()
//////	if err != nil {
//////		fmt.Println(err)
//////	}
//////	log.Println("main goroutine exit")
//////}
//////
//////type HelloHandler struct{}
//////
//////func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//////	time.Sleep(time.Second * 10)
//////	w.WriteHeader(200)
//////	w.Write([]byte("Hello"))
//////}
//////
///////**
//////优雅停止服务
//////*/
//////func graceful(server *http.Server) {
//////	c := make(chan os.Signal, 1)
//////
//////	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
//////
//////	done := make(chan bool, 1)
//////
//////	go func() {
//////		sig := <-c
//////
//////		fmt.Println(sig)
//////
//////		done <- true
//////	}()
//////
//////	fmt.Println("waiting signal")
//////	<-done
//////	fmt.Println("exit")
//////}
//////
///////**
//////控制 goroutine 数量
//////*/
//////func goroutineCount() {
//////	var wg sync.WaitGroup
//////
//////	ch := make(chan struct{}, 10)
//////
//////	for i := 0; i < math.MaxInt64; i++ {
//////		wg.Add(1)
//////
//////		ch <- struct{}{}
//////
//////		go func(i int) {
//////			defer wg.Done()
//////
//////			fmt.Println("i:", i)
//////
//////			time.Sleep(time.Second * 1)
//////
//////			<-ch
//////		}(i)
//////	}
//////
//////	wg.Wait()
//////}
//////
///////**
//////高效字符串拼接
//////*/
//////func stringBuilder() {
//////	var sb strings.Builder
//////
//////	for i := 0; i < 100; i++ {
//////		sb.WriteString(strconv.Itoa(i))
//////		sb.WriteString("-")
//////	}
//////
//////	ret := sb.String()
//////
//////	fmt.Println("ret", ret)
//////}
//////
///////**
//////正则匹配
//////
//////大于等于6，小于等于15长度，首字母是字符，由下划线、数字、字母组成
//////*/
//////func reg() {
//////	// 5 - 14
//////	pattern := "^[a-zA-Z][_0-9a-zA-Z]{5,14}"
//////
//////	//text := "leeyoung"
//////	text := ""
//////
//////	match, _ := regexp.MatchString(pattern, text)
//////
//////	fmt.Println("match:", match)
//////}
////
////package main
////
////import "fmt"
////
////func main() {
////	solution()
////}
////
////func solution() {
////	fmt.Println("hello world")
////}
//
////package main
////
////import "fmt"
////
////func main() {
////	nums := []int{4, 5, 6, 3, 2, 1}
////
////	ret := BubbleSort(nums)
////
////	fmt.Println(ret)
////}
////
////func BubbleSort(array []int) []int {
////	length := len(array)
////
////	swap := true
////
////	for i := 0; i < length; i ++ {
////		if !swap {
////			break
////		}
////
////		swap = false
////
////		for j := 0; j < length-i-1; j ++ {
////			if array[j] > array[j+1] {
////				array[j], array[j+1] = array[j+1], array[j]
////				swap = true
////			}
////		}
////
////		fmt.Println(array)
////	}
////
////	return array
////
////	fmt.Println()
////}
//
////package main
////
////import "fmt"
////
////type Person struct {
////	id   int
////	name string
////}
////
////func main() {
////	a := Person{
////		id:   1,
////		name: "lee",
////	}
////
////	fmt.Printf("%+v", a)
////
////	b := new(map[int]int)
////
////	fmt.Println(b)
////}
//
//package main
//
//import (
//	"errors"
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	//fmt.Println("hello world")
//	//for i := 0; i < 10000; i++ {
//	//	fmt.Println(string(gorand.KRand(8, gorand.KC_RAND_KIND_ALL)))
//	//}
//	//createServer()
//
//	//var a []string
//	//
//	//a = append(a, "1")
//	//a = append(a, "2")
//	//a := []string{"1", "2"}
//	//
//	//b := strings.Join(a, ",")
//	//
//	//fmt.Println(b)
//
//	a := A{
//		a: "1",
//	}
//
//	fmt.Println(reflect.TypeOf(a.b))
//}
//
//func prepareNetwork() (func(), error) {
//	a := 1
//	b := 2
//
//	var sumFunc = func() {
//		fmt.Println(a + b)
//	}
//
//	return sumFunc, nil
//}
//
//func createServer() {
//	f, err := prepareNetwork()
//	if err != nil {
//		return
//	}
//
//	err = doCreate()
//	if err != nil {
//		f()
//	}
//}
//
//func doCreate() error {
//	return errors.New("1")
//}
//
//type B struct {
//	msg string
//}
//
//type A struct {
//	a string
//	b B
//}

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"net/url"
)

func getSignature(appKey, appId, nonce, timestamp, path string) string {
	fmt.Printf("appKey: %s\n", appKey)
	fmt.Printf("appId: %s\n", appId)
	fmt.Printf("nonce: %s\n", nonce)
	fmt.Printf("timestamp: %s\n", timestamp)
	fmt.Printf("path: %s\n", path)
	sig := hmac.New(sha256.New, []byte(appKey))
	sig.Write([]byte(fmt.Sprintf("%s%s%s%s", appId, nonce, timestamp, path)))
	sigResult := hex.EncodeToString(sig.Sum(nil))
	return sigResult
}

func main() {
	appId := "Ib8thze167CZi1pl"
	appKey := "Wqo_tP315pkTK3EaODX!K9PnLGGIS6mI"

	nonce := "A4BCA6A2-4BDF-4114-8823-1A6169C0A7D7"

	timestamp := "1685933311817"

	uri := "https://test-cg.ppio.cloud/openapi/v1/security_groups?regionId=CN-test-office-01"
	path, err := url.Parse(uri)
	if err != nil {
		return
	}
	// path.Path 类似：/openapi/v1/security_groups

	ret := getSignature(appKey, appId, nonce, timestamp, path.Path)

	fmt.Println("ret: ", ret)
}
