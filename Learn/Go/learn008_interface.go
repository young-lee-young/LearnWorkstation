package main

import "fmt"

// 定义接口
type Programmer interface {
	WriteCode() string
}


type GoProgrammer struct {

}

func (p *GoProgrammer) WriteCode() string {
	return "hello world"
}

func main() {
	//var p Programmer
	p := new(GoProgrammer)
	fmt.Println(p.WriteCode())
}
