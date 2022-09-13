# veth



### 实验过程

```bash
# 添加一对veth，添加后两个都在主机的network namespace中
ip link add veth0 type veth peer name veth1

# 将veth1加入到另一个network namespace
ip link set veth1 netns netns1

# 本地的veth0配置IP并拉起
ifconfig veth0 10.1.1.2/24 up

# 将netns1中的veth1配置IP并拉起
ip netns exec netns1 ifconfig veth1 10.1.1.1/24 up

# 查看路由，多了一条路由
route -n
# Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
# 10.1.1.0        0.0.0.0         255.255.255.0   U     0      0        0 veth0

# ping veth1地址，可以通
ping 10.1.1.1
# PING 10.1.1.1 (10.1.1.1) 56(84) bytes of data.
# 64 bytes from 10.1.1.1: icmp_seq=1 ttl=64 time=0.134 ms

# 删除veth，删除一个另一个也将删除
ip link delete veth0
```
