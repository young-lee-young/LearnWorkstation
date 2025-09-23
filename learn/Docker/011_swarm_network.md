# swarm 网络


### 创建 overlay 网络

* 创建网络

```bash
# docker network create -d overlay --attachable --subnet $网段/掩码 网络名
docker network create -d overlay --attachable --subnet 10.209.244.0/28 lee
```


* 查看网络

```bash
docker network ls
```
