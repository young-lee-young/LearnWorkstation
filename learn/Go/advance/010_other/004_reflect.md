# 反射


### 需求

1. 动态获取对象的类型
2. 对任意类型变量（interface）赋值
3. 调用任意方法


### 元数据

元数据就是 "数据的数据"

把对象的类型表示成一个数据类型

把对象的值表示成一个数据类型


### 反射

* 对象的类型

```go
package reflect

type Type interface {
    
}
```


* 对象的值

```go
package reflect

type Value struct {

}
```


* 反射代码

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "lee"
	
	// 反射出类型
	t := reflect.TypeOf(s)
	fmt.Println(t) // string

	// 反射出值
	v := reflect.ValueOf(s)
	fmt.Println(v) // lee
	
	s2 := v.Interface().(string)
	fmt.Println(s2) // lee
}
```


### 使用反射调用方法

```go
package main

import (
	"fmt"
	"reflect"
)

func myAdd(a int, b int) int {
	return a + b
}

func herAdd(a int, b int) int {
	return a - b
}

func callAdd(f func(a int, b int) int) {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return
	}

	argv := make([]reflect.Value, 2)
	argv[0] = reflect.ValueOf(1)
	argv[1] = reflect.ValueOf(2)

	ret := v.Call(argv)
	fmt.Println(ret[0].Int())
}

func main() {
	callAdd(myAdd)
	callAdd(herAdd)
}
```
