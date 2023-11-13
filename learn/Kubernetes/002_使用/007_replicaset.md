# ReplicaSet（副本集）


### 基础命令

* 查看所有 replicaset

```bash
kubectl get replicasets
```


* 查看详情

```bash
kubectl describe replicaset $replicaset_name
```


* 删除

```bash
# 会删除所管理的 pod
kubectl delete replicaset $replicaset_name
```
