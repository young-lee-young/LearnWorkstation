package main

import (
	"fmt"
	"time"
)

type Person struct {
	salary int
	level  int
}

func (p *Person) promote() {
	p.salary += 1
	p.level += 1
}

func main() {
	p := Person{salary: 0, level: 0}

	for i := 0; i < 200; i++ {
		go p.promote()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(p.salary, p.level)
}
