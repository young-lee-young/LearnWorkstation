# network namespace


### network namesapce操作

```bash
# 添加netns
ip netns add $netns_name

# 查看netns
ip netns list

# 删除netns
ip netns delete $netns_name

# 在netns上执行命令
ip netns exec $netns_name 命令
```


### 实验过程

```bash
ip netns add netns1

ip netns list
# netns1 (id: 0)

# ping本地环回口，不通
ip netns exec netns1 ping 127.0.0.1
# connect: Network is unreachable

# 将本地环回口up
ip netns exec netns1 ip link set dev lo up

# 再次ping本地环回口，已通
ip netns exec netns1 ping 127.0.0.1
# PING 127.0.0.1 (127.0.0.1) 56(84) bytes of data.
# 64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.092 ms
```
