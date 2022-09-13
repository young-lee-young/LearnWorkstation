# 虚拟网桥



### iproute2包操作网桥

```bash
# 创建网桥
ip link add name br0 type bridge

# up网桥
ip link set br0 up

# 查看网桥
ip link list type bridge

# 将veth pair设备加入到网桥中
ip link set dev veth0 master br0

# 查看网桥连接网络设备
bridge link
```



### bridge-utils包操作网桥

```bash
# 安装操作网桥包
yum install bridge-utils

# 创建网桥
brctl addbr br0

# 查看网桥
brctl show

# 删除网桥
brctl delbr br0

# 将veth pair网卡加入网桥
brctl addif br0 veth0

# 查看网桥连接网络设备
brctl show

# 将veth pair网卡移出网桥
brctl delif br0 veth0
```
