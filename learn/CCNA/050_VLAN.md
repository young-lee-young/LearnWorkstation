# 虚拟局域网 VLAN


### VLAN基础

VLAN 是交换机组网才有的概念

一个 VLAN = 一个广播域 = 一个逻辑子网网段

* VLAN优点

1. 分段
2. 灵活
3. 安全


### 802.1 Q 干道协议（Dot One Q）

![VLAN帧](image/VLAN帧.png)

VLAN帧：目的MAC地址，源MAC地址，VLAN标记，类型，数据，冗余校验值

VLAN 标记的**最后 12 比特**称为 VLAN标识符VID，标识以太网帧属于哪一个VLAN，取值范围 0 - (2^12 - 1)，0 和 4095 不用来表示 VLAN，有效取值范围 0 - 4094

802.1 Q 帧是由交换机处理，当交换机收到普通以太网帧，会将其插入**4 字节的 VLAN 标记**转变为 802.1 Q帧，简称"打标签"

* Trunk口转发规则

发送规则：对于 VID 等于 PVID 的帧，**去标签**再转发；VID 不等于 PVID 的帧，直接转发

接收规则：接收未打标签的帧，根据 PVID **打标签**；接收**已打标签**的帧


### VTP（VLAN中继协议）


### VLAN间路由（单臂路由）

1. 数据从PC发出，进入到交换机接口
2. 因为交换机接口指定了VLAN，所以将数据帧打tag
3. 交换机出口配置成Trunk口，链路是干道链路，与路由器物理口相连，
4. 路由器物理口配置子接口（逻辑接口），写入vlan_id，写入IP地址（既VLAN的网关地址）

* 三层交换机（具有数据交换模块和路由模块的交换机）

```bash
# 交换机开启路由功能
ip routing

# 创建 VLAN
vlan 10

# 为 VLAN 添加IP地址
interface vlan 10
ip address 192.168.1.1 255.255.255.0

# 给连接 PC 一端口设置为 access 口
interface f0/2
switchport mode access

# 给出口设置为 trunk 口
interface f0/1
switchport trunk encapsulation dot1q
switchport mode trunk
```
