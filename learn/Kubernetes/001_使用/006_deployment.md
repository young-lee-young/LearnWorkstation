# Deployment


### 基础

* 作用

Deployment 可以理解为由两部分组成，其中的 template 其实就是定义 Pod，replicas 定义需要的副本数，Deployment Controller 保证 Pod 数量一直为副本数


### 基本命令

* 创建 deployment

```bash
kubectl apply -f $yaml_file_path
```

* 查看所有 deployment

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
kubectl edit deployment.v1.apps/$deployment_name
```

* 查看上线状态

```bash
kubectl rollout status deployment/$deployment_name
```
