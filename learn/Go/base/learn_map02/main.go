package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/**
	map并发读写会panic
	fatal error: concurrent map read and map write
 */

func concurrentMap() {
	var wg sync.WaitGroup
	m := make(map[string]int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i ++ {
			fmt.Println("set", i)
			m[strconv.Itoa(i)] = i
		}
	}()

	go func() {
		wg.Done()
		for i := 0; i < 10000; i ++ {
			fmt.Println("get", m[strconv.Itoa(i)])
		}
	}()

	wg.Wait()
}

// 使用sync.Map解决map并发不安全问题
func concurrentMapSolution() {
	var wg sync.WaitGroup
	var syncM sync.Map

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i ++ {
			fmt.Println("set", i)
			syncM.Store(strconv.Itoa(i), i)
		}
	}()

	time.Sleep(time.Second * 1)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i ++ {
			v, _ := syncM.Load(strconv.Itoa(i))
			fmt.Println("get", v)
		}
	}()

	wg.Wait()
}

func main() {
	//concurrentMap()
	concurrentMapSolution()
}
