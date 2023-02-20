# TODO 

1. 证书加密 和 TLS 握手
2. git merge 和 git rebase
3. 表连接

B 站 MySQL 看完
B 站 Docker 看完 


### 单位换算

```
带宽换算单位K：1000，M：1000 * 1000，G：1000 * 1000 * 1000

内存、流量换算单位K(Kilo)：1024，M(Mega)：1024 * 1024，G：1024 * 1024 * 1024
```


### linux目录结构

* 内核参数

/etc/sysctl.conf

* 全局环境变量配置路径

/etc/profile

* 系统启动方式

/etc/inittab

* 系统启动脚本

/etc/init.d


### Redis

1. 五种类型数据底层实现
2. redis事物
3. redis过期策略
4. 主从架构哨兵问题
5. 线程模型
6. pipeline


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
