# DHCP（动态主机配置协议，Dynamic Host Configuration Protocol）


### 基础

DHCP 发现报文（DHCP DISCOVER）：客户端发送的发现报文

DHCP 提供报文（DHCP OFFER）：服务端发送的响应报文

DHCP 请求报文（DHCP REQUEST）：客户端发送的请求报文

DHCP 确认报文（DHCP ACK）：服务端发送的确认报文

DHCP 拒绝报文（DHCP DECLINE）：客户端发送的拒绝报文

DHCP 否认报文（DHCP NAK）：服务端发送的否认报文

DHCP 释放报文（DHCP RELEASE）：客户端发送的释放报文


### 过程

* 客户端发送发现报文

1. 服务器使用端口 67，客户端使用端口 68
2. 客户端广播发送 "DHCP DISCOVER"，源 IP 地址为 0.0.0.0（未分配到IP），目的 IP 地址为广播地址 255.255.255.255（主机不知道 DHCP 服务器地址）
3. 网络中所有设备都会收到 "DHCP DISCOVER" 并解析；其他设备将其丢弃，DHCP 服务器发现是 "DHCP DISCOVER"，进行响应

* 服务端响应提供报文

1. DHCP 服务器收到 "DHCP DISCOVER" 后，根据封装的 MAC 地址查询数据库，查看是否有针对该 MAC 地址的配置信息；如果有，使用配置信息构建并发送 "DHCP OFFER"，如果没有，采用默认配置信息构建并发送 "DHCP OFFER"
2. DHCP 服务器响应 "DHCP OFFER"，源 IP 地址为 DHCP 服务器 IP 地址，目的 IP 地址为广播地址 255.255.255.255（目的主机还没有配置 IP 地址，为了使其收到，只能发送广播）
3. 网络中所有设备都会收到 "DHCP OFFER" 并解析；其他设备将其丢弃，目的主机运行着监听 68 端口的应用进程，接收 "DHCP OFFER" 并处理
4. 客户端根据 "DHCP OFFER" 中的事务 ID，判断该报文是否是自己所请求的报文；如果不是，丢弃报文，如果是，处理报文

* 客户端发送请求报文

1. 客户端向所选择 DHCP 服务器发送 "DHCP REQUEST"，源 IP 地址为 0.0.0.0，目的 IP 地址为广播地址 255.255.255.255（不用向所有 DHCP 服务器发送单播报文，来告知每个服务器是否请求它们作为自己的 DHCP 服务器）

* 服务端响应确认报文

1. 服务端响应 "DHCP ACK"，源 IP 地址为 DHCP 服务器 IP 地址，目的 IP 地址为广播地址 255.255.255.255
2. 客户端使用 ARP 检测所分配到 IP 地址是否已被网络中其他主机占用；如果被占用，给 DHCP 服务器发送 "DHCP DECLINE" 撤销 IP 地址租约，重新发送 "DHCP DISCOVER"，如果未占用，使用此 IP 地址进行通信

* 续租

1. 过了 0.5 倍租期，向 DHCP 服务器发送 "DHCP REQUEST"，源 IP 地址为当前 IP 地址，目的 IP 地址为服务器 IP 地址
2. DHCP 服务器同意，发送 "DHCP ACK"，不同意发送 "DHCP NAK"
3. 客户端收到 "DHCP NAK"，停止租用 IP 地址，并重新发送 "DHCP DISCOVER" 报文，重新申请 IP 地址
4. 如果 0.5 倍租期时，DHCP 服务器未响应，则在 0.875时再次发送 "DHCP REQUEST"
5. 如果 DHCP 服务器一直未响应，则到期后客户端立即停止使用 IP，并重新发送 "DHCP DISCOVER"

* 退租

1. 客户端可随时退租，向 DHCP 服务器发送 "DHCP RELEASE"，源 IP 地址为 0.0.0.0，目的 IP 地址为广播地址 255.255.255.255


### DHCP 报文格式

* 发现报文

1. 事务ID
2. DHCP 客户端的 MAC 地址

* 提供报文

1. 配置信息：IP 地址、子网掩码、地址租期、默认网关、DNS 服务器

* 请求报文

1. 事务ID
2. DHCP 客户端的 MAC 地址
3. 接受的租约中的IP 地址
4. 提供此租约的 DHCP 服务端的 IP 地址


### 注意⚠️

* DHCP 服务器会使用 ARP 确保所选 IP 地址未被网络中其他主机占用

* 如果有多个 DHCP 服务器响应多个 "DHCP OFFER"，通常选最先到达的


### DHCP 中继代理

给路由器配置 DHCP 服务器的 IP 地址，并使之成为 DHCP 中继代理
