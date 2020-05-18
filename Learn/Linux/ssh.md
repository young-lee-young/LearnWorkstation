# ssh命令使用


### 开启连接

```bash
# 安装
sudo apt-get install openssh-server

# 启动
/etc/init.d/ssh start

# 重启
/etc/init.d/ssh restart

# 检测
ps -e | grep ssh
```


### 配置文件

* 位置

/etc/ssh/sshd_config


### 连接

```bash
ssh 用户名@主机地址 -p 端口
# 输入密码
```


### 免密登录

id_rsa私钥
id_rsa.pub公钥

```bash
ssh-copy-id -i ～/.ssh/id_rsa.pub 用户名@主机地址

# 修改以下权限
chmod 700 .ssh
chmod 600 id_rsa
chmod 644 id_rsa.pub
chmod 600 authorized_keys
```
