package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu     sync.RWMutex
	salary int
	level  int
}

func (p *Person) promote() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.salary += 1000
	fmt.Println(p.salary)
	p.level += 1
	fmt.Println(p.level)
}

func (p *Person) print() {
	p.mu.RLock()
	defer p.mu.RUnlock()

	fmt.Println(p.salary)
	fmt.Println(p.level)
}

func main() {
	p := Person{salary: 10000, level: 0}
	once := sync.Once{}
	go once.Do(p.promote)
	go once.Do(p.promote)
	go once.Do(p.promote)
	//go p.promote()
	//go p.promote()
	//go p.promote()
	time.Sleep(time.Second)
}
