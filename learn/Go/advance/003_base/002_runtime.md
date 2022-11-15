### Runtime

* 什么是 Runtime

Runtime 就是程序的运行环境


* 能力

1. 内存管理能力
2. 垃圾回收能力（GC）
3. 超强的并发能力（协程调度）


* Go Runtime 特点

1. 没有虚拟机概念
2. Runtime 作为程序的一部分打包进二进制文件
3. Runtime 随用户程序一起运行
4. Runtime 与用户程序没有明显界限，直接通过函数调用


* 其他特点

1. 有一定屏蔽系统调用能力
2. Go 的一些关键字是 Runtime 下的函数（如 'go关键字' 在编译时转为对 runtime 包 newproc 的调用）
