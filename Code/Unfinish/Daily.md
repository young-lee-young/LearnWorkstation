Django 用户与权限认证，Backend学习
Django 日志
Django Swagger


### 近期学习计划

2. Shell脚本（完成）
6. SVN（完成）

3. MySQL数据库进阶（）
5. Docker进阶（完成）
7. Nginx（）
8. Apache
9. Flask
4. Git进阶（）


### 慕课

* 云计算&大数据（）

1. Docker入门

* 运维&测试
1. Jenkins
2. Python自动化运维篇


### 问题

广播风暴：为什么不用交换机发送过这个广播端口就不再发送这个广播

冲突域：collision domain

show | match 8.8.8.8 | display set


# 单位换算

```
带宽换算单位K：1000，M：1000 * 1000，G：1000 * 1000 * 1000

内存、流量换算单位K(Kilo)：1024，M(Mega)：1024 * 1024，G：1024 * 1024 * 1024
```


### C端覆盖

* 服务器部署

```sh
# harbor镜像推送
docker tag cping:v1 harbor.capitalonline.net/dev/cping:v1
docker push harbor.capitalonline.net/dev/cping:v1
```


### linux目录结构

* 网卡配置

/etc/sysconfig/network-scripts/ifcfg-eth0

* HOSTNAME

/etc/sysconfig/network

* DNS解析

/etc/resolv.conf

* host文件

/etc/hosts

* 密码

/etc/passwd

* 内核参数

/etc/sysctl.conf

* 全局环境变量配置路径

/etc/profile

* 开机自动挂载

/etc/fstab

* 系统启动方式

/etc/inittab

* 系统启动脚本

/etc/init.d



### 日志文件

/var/log/message

/var/log/secure


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
优先队列


### 协程

futrue：

_loop：
_state：


task（继承Futrue）

_step：执行的回调方法


loop（BaseEventLoop）:

_ready是一个双端队列，存events.Handle(callback)对象

_task_factory



### 数据结构

* 复杂度分析

渐近时间复杂度（n趋近于正无穷时时间复杂度）
均摊时间复杂度
复杂度震荡
