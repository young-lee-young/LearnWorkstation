package main

import (
	"fmt"
	"unsafe"
)

// 定义结构
type User struct {
	Id string
	Name string
	Age int
}

// 添加方法，调用方和方法是不同的对象
func (user User) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&user.Name))
	return fmt.Sprintf("ID: %s, Name: %s, Age: %d", user.Id, user.Name, user.Age)
}

// 添加方法，调用方和方法是相同的对象
func (user *User) StringPoint() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&user.Name))
	return fmt.Sprintf("ID: %s, Name: %s, Age: %d", user.Id, user.Name, user.Age)
}

func main() {
	// 创建实例
	user1 := User{"1", "leeyoung", 10}
	user2 := &User{Name: "zhaomeng", Age: 20}

	// new返回的是指针
	user3 := new(User)
	user3.Id = "3"
	user3.Name = "xuwei"
	user3.Age = 30

	//通过实例访问方法
	fmt.Printf("Address is %x\n", unsafe.Pointer(&user1.Name))
	fmt.Println(user1.String())

	// 通过指针访问方法
	fmt.Printf("Address is %x\n", unsafe.Pointer(&user2.Name))
	fmt.Println(user2.StringPoint())
}
