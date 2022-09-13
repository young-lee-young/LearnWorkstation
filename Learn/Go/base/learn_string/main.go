package main

import "fmt"

/**
	字符串的本质是[]byte，是一个slice
 */
func main()  {
	s1 := []byte("Hello")
	fmt.Println(s1)
	s2 := append(s1, "World"...)
	fmt.Println(s2)
}
