g# FrameRelay（帧中继协议）


### 基础

使用虚电路（VC，Virtual Circuit）进行连接
提供面向对象的服务


### 术语

* DLCI（数据链路连接标识，Data Link Connetction Identifier）

1. 通常由帧中继服务提供商分配
2. 帧中继DLCI仅具有本地意义
3. DLCI 0-15和1008-1023留作特殊用途，服务提供商分配的DLCI范围通常为16-1007

* LMI（本地管理接口）

管理本地接口到帧中继提供商之间的这段链路，路由器从帧中继交换机接口接收LMI信息，并将虚电路状态更新为3种状态（Active state（正常状态）、Inactive state（远程路由器没有工作）、Deleted state（接口没有收到交换机的任何LMI信息，可能是映射问题或者线路问题））之一

1. 是一种信令标准，用于管理链路连接及keepalive机制
2. 终端路由器（DTE）和帧中继交换机（DCE）之间的帧中继设备每10秒轮询一次网络
3. Cisco路由器支持三种：Cisco、Ansi、q933a


### 帧中继地址映射

* 格式

Frame Relay | DLCI | 远程IP地址

* 获取方式

1. 帧中继映射条目，DLCI从运营商处获取，映射关系为远端IP地址到本地的DLCI之间的联系
2. 通过手工配置或Inverse ARP自动发现


### 帧中继配置

* 帧中继交换机

```sh
# 将路由器模拟成帧中继交换机
frame-relay switching

# 接口模式，启用帧中继
encapsulation frame-relay
# 设置始终频率
clock rate 64000
# 将接口配置为DEC
frame-relay intf-type dce
# 设置VC
frame-relay route 当前接口连接的路由器的DLCI 出接口 目的DLCI
```

* 帧中继路由器配置

```sh
# 接口模式，启用帧中继
encapsulation frame-relay
# 关闭inverse-arp（反向ARP）
no frame-relay inverse-arp
# 手工配置帧中继映射，broadcast可选，伪广播
frame-relay map ip 目的网络 对应的DLCI broadcast
```

* 帧中继交换机查看映射条目

```sh
show frame-relay route
```

* 帧中继路由器查看映射条目

```sh
show frame-relay map
```


### 帧中继环境中的动态路由协议问题

非广播型多路访问（不支持广播），动态路由协议需要广播或组播发送协议数据，因此启用帧中继的路由器不能使用动态路由协议，以下为解决方法

1. 帧中继虽然不支持广播，但是可以模式广播的操作，做法是通过向所有的PVC发送一份数据的拷贝
2. 在建立PVC的时候，通过inverse-arp自动建立映射，默认开启上述功能，如果是手工配置映射，必须加上broadcast关键字
