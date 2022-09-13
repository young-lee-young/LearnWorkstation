package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

/**
	接口

	接口变量里有
	实现着类型和实现者指针，指向实现者
 */

const url = "http://www.baidu.com"
// 接口定义
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{"language": "goland"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another fake baidu content"})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	// 接口变量
	var r Retriever

	retriever := mock.Retriever{"this is a fake baidu content"}
	r = &retriever
	inspect(r)
	//fmt.Println(download(r))

	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)
	//fmt.Println(download(r))

	// type assertion
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println(session(&retriever))
}
