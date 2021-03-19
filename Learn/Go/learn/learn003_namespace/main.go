/**
 * @Time:    2021/3/8 14:24 
 * @Author:  leeyoung
 * @File:    main.go
 * #Content: namespace
 */
package main

import (
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{

	}
}
