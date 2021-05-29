## 近期学习计划

3. MySQL数据库进阶（）
5. Docker进阶（完成）
7. Nginx（）
8. Apache
4. Git进阶（）




### 慕课

* 云计算&大数据（）

1. Docker入门

* 运维&测试
1. Jenkins
2. Python自动化运维篇




### 问题

广播风暴：二层有环，广播包泛洪

广播风暴：为什么不用交换机发送过这个广播端口就不再发送这个广播（并不能判断是否是同一个广播包）

冲突域：collision domain

show | match 8.8.8.8 | display set




### 单位换算

```
带宽换算单位K：1000，M：1000 * 1000，G：1000 * 1000 * 1000

内存、流量换算单位K(Kilo)：1024，M(Mega)：1024 * 1024，G：1024 * 1024 * 1024
```



### 网络

网络传输过程中，IP地址是不变的，mac地址会变化

TCP数据包限制为64KB，65535字节

MTU：1500字节

超时重传：



### linux目录结构

* 内核参数

/etc/sysctl.conf

* 全局环境变量配置路径

/etc/profile

* 系统启动方式

/etc/inittab

* 系统启动脚本

/etc/init.d




redis

1. 五种类型数据底层实现
2. redis事物
3. redis过期策略
4. 主从架构哨兵问题
5. 线程模型
6. pipeline

协程
异步（asyncio）、并发库（gevent，猴子补丁）


垃圾回收
linux命令
僵尸进程和孤儿进程

二分搜索树前、中、后序遍历非递归实现



### 数据结构复杂度分析

渐近时间复杂度（n趋近于正无穷时时间复杂度）
复杂度震荡



HDD（Hard Disk Drive，硬盘驱动器） -> SSD（Solid State Drive，固态驱动器） -> rSSD

RDMA（Remote Direct Memory Access，远程直接数据存取）：两台计算机通信，一个主机内存直接访问另一个主机内存


### Ubuntu20.04安装X11

```bash
# 安装
sudo apt-get install x11vnc

# 设置密码
x11vnc -storepasswd

# 启动
x11vnc -forever -shared -rfbauth ~/.vnc/passwd
```


### Ubuntu20.04安装VMware

```bash
# 添加可执行权限
sudo chmod +x VMware-Workstation-Full-16.1.1-17801498.x86_64.bundle

# 安装
sudo ./VMware-Workstation-Full-16.1.1-17801498.x86_64.bundle

# vmware启动所需包
sudo apt-get install gcc
sudo apt-get install build-essential

# 安装wireshark
sudo apt-get install wireshark
```

TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=static
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEM_MODE=stable-privacy
NAME=ens33
UUID=560df6da-b1ce-4b76-84cb-2064bd9c5c1e
DEVICE=ens33
ONBOOT=yes
IPADDR=192.168.1.10
GATEWAY=192.168.1.10
NETMASK=255.255.255.0
DNS1=114.114.114.114

配置：
master：4核8G 磁盘200G
192.168.1.10

worker1：16核32G 磁盘200G
192.168.1.20

worker2：16核32G 磁盘200G
192.168.1.30

```bash
# 安装docker
yum install docker

# 关闭selinux
vim /etc/selinux/config
# SELINUX=disabled
```
