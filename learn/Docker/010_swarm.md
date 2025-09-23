# swarm 集群


### 创建集群

* 初始化 manager

```bash
docker swarm init --advertise-addr $manager_ip
```

* 查看 manager 加入集群命令

```bash
docker swarm join-token manager
```

* 其他 manager 加入集群

```bash
# 执行上面命令的结果
```


### 加入集群

* 查看加入集群命令

```bash
docker swarm join-token worker
```

* 节点加入集群

```bash
# 执行上面命令的结果
```


### 查看节点

```bash
# 查看集群所有节点
docker node ls
```

### 退出集群

* 删除节点

```bash
# 在节点上执行，状态会变为 Down 状态
docker swarm leave

# manager 上删除节点
docker node rm $node_id
```

* 删除 manager

```bash
docker swarm leave --force
```
