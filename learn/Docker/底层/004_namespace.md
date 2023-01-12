# Namespace


### 基础

* 作用

对资源进行隔离，使得进程在单独的命名空间中运行，且只可以访问当前命名空间的资源


* 分类

pid namespace：进程 ID 隔离

net namespace：网络隔离，在 namespace 中，用户可以拥有独立的 IP、路由、端口

mnt namespace：文件系统挂载点隔离

ipc namespace：信号量、消息队列和共享内存隔离

uts namespace：主机名和域名隔离
