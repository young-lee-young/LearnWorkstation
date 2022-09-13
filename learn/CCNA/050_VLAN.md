# VLAN


### VLAN基础

VLAN是交换机组网才有的概念
一个VLAN = 一个广播域 = 一个逻辑子网网段

* VLAN优点

1. 分段
2. 灵活
3. 安全

* 


### 802.1Q干道协议（dot1q）

数据帧：目的MAC地址，源MAC地址，数据长度，数据，校验值
带有VLAN TAG数据帧：目的MAC地址，源MAC地址，TAG，数据长度，数据，校验值


### VTP（VLAN中继协议）




### VLAN间路由（单臂路由）

1. 数据从PC发出，进入到交换机接口
2. 因为交换机接口指定了VLAN，所以将数据帧打tag
3. 交换机出口配置成Trunk口，链路是干道链路，与路由器物理口相连，
4. 路由器物理口配置子接口（逻辑接口），写入vlan_id，写入IP地址（既VLAN的网关地址）

* 三层交换机（具有数据交换模块和路由模块的交换机）

```sh
# 交换机开启路由功能
ip routing
# 创建VLAN
vlan 10
# 为VLAN添加IP地址
interface vlan 10
ip address 192.168.1.1 255.255.255.0
# 给连接PC一段口设置为access口
interface f0/2
switchport mode access
# 给出口设置为trunk口
interface f0/1
switchport trunk encapsulation dot1q
switchport mode trunk
```
