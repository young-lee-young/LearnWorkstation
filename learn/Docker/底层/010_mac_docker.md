# MacOS Docker 研究


* Mac Docker 宿主机

```bash
# 进入虚拟机
screen /Users/leeyoung/Library/Containers/com.docker.docker/Data/vms/0/tty

# 登出虚拟机
同时按住ctrl + a，松开再按d

# 再次进入
screen -r

# 彻底退出
同时按住ctrl + a，松开再按k
```


### Docker for mac 网络

com.docker.hyperkit 进程来启动虚拟机
