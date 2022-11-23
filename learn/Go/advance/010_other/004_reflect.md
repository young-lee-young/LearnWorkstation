### 需求

1. 获取对象的类型
2. 对任意类型变量赋值
3. 调用任意方法


### 反射

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
	fmt.Println(t)

	// 反射出值
	v := reflect.ValueOf(s)
	fmt.Println(v)
	
	s2 := v.Interface().(string)
	fmt.Println(s2)
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
