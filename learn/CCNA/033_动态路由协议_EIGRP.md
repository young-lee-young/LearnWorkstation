### EIGRP路由协议（加强型内部网关路由协议）


通配符掩码：（0表示匹配，1表示任意）

### EIGRP基本

1. Cisco私有协议，根据带宽和延时选择路径



### 报文

1. 


### EIGRP计算公式


### DUAL算法

* 特点

1. 无环拓扑
2. 可立即使用的无环备用路径
3. 快速收敛
4. 低带宽利用率（通过限定更新实现）

* 术语

1. 后继路由器（Successors）
2. 可行距离（FD，Feasible distance）
3. 可行后继路由器（FS，Feasible Successor）
4. 通告距离（AD）
5. 可行条件（FC，Feasible Condition）


### 配置EIGRP路由协议

```sh
# 特权模式

# 启动ergrp协议
router eigrp 编号

# 指定参与动态路由的网段
network 网段
```
