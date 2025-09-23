# Deployment


### 基础

* 作用

Deployment 可以理解为由两部分组成，其中的 template 其实就是定义 Pod，replicas 定义需要的副本数，Deployment Controller 保证 Pod 数量一直为副本数


### 基本命令

* 创建 deployment

```bash
kubectl apply -f $yaml_file_path
```

* 查看所有的 deployment

```bash
kubectl get deployments
```

* 查看 deployment

```bash
kubectl get deployment $deployment_name
```

* 查看 deployment 详情

```bash
kubectl describe deployment $deployment_name
```

* 删除 deployment

```bash
kubectl delete deployment $deployment_name
```


### 高级命令

* 更新 deployment

```bash
kubectl edit deployment $deployment_name
```


* 查看上线状态

```bash
kubectl rollout status deployment $deployment_name
```


* 回滚

```bash
# 查看历史版本
kubectl rollout history deployment $deployment_name

# 查看某个版本的 API 对像细节
kubectl rollout history deployment $deployment_name --revision=1

# 回滚到上一个版本
kubectl rollout undo deployment $deployment_name

# 回滚到指定版本
kubectl rollout undo deployment $deployment_name --to-revision=1
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: xxx
spec:
  # 控制历史版本数量，即 ReplicaSet 数量
  revisionHistoryLimit: 5
```


* 挂起和恢复

```bash
# 1. 挂起 Deployment
kubectl rollout pause deployment $deployment_name

# 2. 执行操作，此时不会执行更新，也不会创建 ReplicaSet

# 3. 恢复 Deployment
# 此时才会进行更新，只会创建一个 ReplicaSet
kubectl rollout resume deployment $deployment_name
```


### 滚动更新

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: xxx
spec:
  # 滚动更新
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
```

* type：RollingUpdate 或 Recreate


* maxSurge：最大可以超过 DESIRED 的数量，可以是数值或百分比，默认 25%，向上取整

如 DESIRED 为 3，默认 25% 情况下，最大 Pod 数量 3 + 3 * 0.25 = 3.75，因此最大可以是 4 个 Pod


* maxUnavailable: 最大不可用 Pod 数量，可以是百分比，默认 1

如 DESIRED 为 3，默认 1 情况下，最少正常 Pod 数量 3 - 1 = 2，因此最少要有 2 个 Pod

maxSurge 和 maxUnavailable 不可以同时为 0，同时为 0 不能进行滚动升级
