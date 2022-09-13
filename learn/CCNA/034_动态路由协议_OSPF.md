# OSPF协议（Open Shortest Path First，开放最短路径优先协议）


### 简介

1. 链路状态协议，在RFC2328中描述
2. 公有协议


### OSPF术语

	
* COST：OSPF使用开销作为路由度量值（Metric）

接口的cost为10^8/接口带宽(b/s)，10^8是参考带宽值，可以修改

* DR：指定路由器，DR具有非抢占性

* BDR：备用指定路由器

* DROther：除指定路由器和备用指定路由器外的其他路由器


### OSPF网络类型

1. 广播型多路访问
2. 点对点
3. 非广播型多路访问
4. 点对多点


### OSPF三张表

* 邻居表

* 链路状态数据库（LSDB）

* 路由表


### OSPF报文

* Hello包：建立和维护OSPF邻居关系
* DBD：链路状态数据库描述信息（描述LSDB中的LSA头部信息）
* LSR：链路状态请求，向OSPF邻居请求链路状态信息
* LSU：链路状态更新，答复LSR（包含一条或多条LSA，link state advertisement，链路状态通告）
* LSAck：对LSU中的LSA进行确认


### DR选举

1. 具有高优先级的接口选为DR
2. 如果接口优先级相同，具有高Router-ID的路由器被选举为DR
3. 如果没有设置Router-ID，则接口IP地址越大越优先


### OSPF协议过程

1. 定时组播（224.0.0.5）发送hello报文，接受到邻居的hello包，邻居状态达到2-Way状态，选举DR、BDR，与DR、BDR建立邻接关系，状态达到Full状态
2. LSAs（链路状态通告）的泛洪（将消息以组播形式发送到224.0.0.6）发送到DR、BDR，DR、BDR得到消息后泛洪（将消息以组播的形式发送到224.0.0.5），将信息存储到LSDB（Link State Database）中
3. 根据LSDB，使用SPF（最短路径算法）进行计算，得到一个以自己为根，覆盖全网的一棵无环的树
4. 根据计算结果，将最优路径存进路由表


### 路由器角色

1. 区域内路由器（Internal Router，非骨干区域路由器）
2. 区域边界路由器（Area Border Router，连接骨干区域和非骨干区域路由器）
3. 骨干路由器（Backbone Router，骨干区域中路由器）
4. AS边界路由器（AS Boundary Router，连接到其他AS路由器）


### LSA报文

* 报文组成

IP包：IP报头（IP协议号89） + OSPF报文
OSPF报文：OSPF报头 + OSPF报文数据
LSU：LSU头部 + LSA + LSA + LSA
LSA：LSA头部 + LSA数据

* LSA分类

1. Router LSA：每台OSPF路由器都会产生，描述路由器所有的OSPF直连链路的状态和Cost值，只能在所属区域内泛洪
2. Network LSA：DR产生，描述DR自己及DR连接路由器的RouteID，只能在所属区域内泛洪
3. Network Summary LSA：由ABR产生，


### SPF（最短路径算法）



### OSPF配置

* 配置

```sh
# 开启ospf进程
router ospf 进程id

# 宣告地址
network 网络号 通配符掩码 area 区域id
```

* 查看

```sh
show ip ospf neighbor
```
