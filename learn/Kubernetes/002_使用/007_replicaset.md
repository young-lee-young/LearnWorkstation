# ReplicaSet（副本集）


### 基础命令

* 查看

```bash
kubectl get rs
```

* 查看详情

```bash
kubectl describe rs $rs_name
```

* 删除

```bash
# 会删除所管理的 pod
kubectl delete rs $rs_name
```
