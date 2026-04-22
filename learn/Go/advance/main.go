package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fileName := strconv.Itoa(int(time.Now().UnixMilli()))
	fmt.Println(fileName)
}
