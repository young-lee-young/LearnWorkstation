# Ansible基础

### 自动化管理IT资源的工具


### 功能

+ 系统环境配置
+ 安装软件
+ 持续集成
+ 热回滚


### 安裝

+ pip安装

```sh
pip install ansible
```

+ 源码安装

```sh
# 获取源码，https://github.com/ansible/ansible
# 解压源码
# 进入源码目录
source ./hacking/env-setup
```

+ 系统源安装

```
# Ubuntu下源码安装
apt-get install soft-ware-properties-common
apt-add-repository ppa:ansible/ansible
apt-get update
apt-get install ansible
```


### 配置文件

+ 配置文件优先级

1. export ANSIBLE_CONFIG（导入的配置文件）
2. ./ansible.cfg（当前目录下的配置文件）
3. ~/.ansible.cfg（用户目录下的配置文件）
4. /etc/ansible/ansible.cfg

+ 配置文件获取

[GitHub上ansible配置文件](http://raw.github.com/ansible/ansible/devel/examples/ansible.cfg)
