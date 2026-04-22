# Linux网络


### 命令

```bash
# 网络设备文件
/sys/class/net

# 当前命名空间网络设备
/proc/self/net

# 某个进程namespace
/proc/$PID/ns
```


### vlan

```bash
# 查看vlan模块
lsmod | grep vlan

# 加载vlan模块
modprobe 8021q

# 移除vlan模块
modprobe -r 8021q
```
