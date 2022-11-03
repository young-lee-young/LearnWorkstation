### NFS 安装

* 服务端安装

```bash
# 安装 NFS
sudo apt-get update
sudo apt-get install nfs-kernel-server


sudo mkdir -p /home/lee/nfs


# 修改目录权限
sudo chmod -R 777 /home/lee/nfs
# 修改用户
sudo chown -R nobody:nogroup /home/lee/nfs


# 修改，内容如下
sudo vim /etc/exports


# 重启服务
sudo systemctl restart nfs-kernel-server.service
```


* 配置

```conf
/home/lee/nfs *(rw,sync,no_root_squash)
```


* 客户端安装

```bash
sudo apt-get update

sudo apt-get install nfs-common
```
