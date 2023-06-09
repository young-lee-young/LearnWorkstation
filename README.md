### TODO 

强制缓存和协商缓存

Redis 底层数据结构和日志


2. git merge 和 git rebase

快排

topK

数组中出现次数超过一半的数字

内存管理

文件管理


### 学习计划

1. 慕课网 Moody 老师 RabbitMQ
2. MIT 6.824 课程
3. Go 1.18 泛型
4. 代码随想录刷题
5. vim 学习


### 单位换算

```text
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


### 数据结构复杂度分析

二分搜索树前、中、后序遍历非递归实现

渐近时间复杂度（n 趋近于正无穷时时间复杂度）

复杂度震荡


HDD（Hard Disk Drive，硬盘驱动器） -> SSD（Solid State Drive，固态驱动器） -> rSSD

RDMA（Remote Direct Memory Access，远程直接数据存取）：两台计算机通信，一个主机内存直接访问另一个主机内存


### Ubuntu 20.04 安装 X11

```bash
# 安装
sudo apt-get install x11vnc

# 设置密码
x11vnc -storepasswd

# 启动
x11vnc -forever -shared -rfbauth ~/.vnc/passwd
```


AI 图片生成

Stable Diffusion

Stable Diffusion 1：使用 OpenAI 的 CLIP

Stable Diffusion 2：使用 OpenCLIP（CLIP 开源版本）


三部分组成：变分自编码器（VAE）、U-Net、文本编码器

VAE：负责图像处理
U-Net：预测噪音
文本编码器：处理文本


* 图片处理分成 2 部分

编码过程：提供图片素材，进行训练

解码过程：文生图、图生图


* 模型

基础模型：

微调模型：


Checkpoint：大模型

Dreambooth：大模型，预调模型，深度学习模型，微调现有文生图模型

Lora：产出单个特征的图像，和大模型一起使用

Embedding：本文反转，几张图片 -> 关键词，和大模型一起使用

Hypernetwork：用的少
