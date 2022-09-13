# deployment



### 基础

* 作用

  deployment可以理解为由两部分组成，其中的template其实就是定义Pod，replicas定义需要的状态，Deployment Controller保证pod数量等一直满足需要的状态。

  Pod是k8s的最小调度单位，一个Pod中可以有多个containers，彼此共享网络等，这是k8s的核心概念。Deployment，StatefulSet，Job是Controller，保证Pod一直运行在你需要的状态。

  

* 工作负载资源
  1. Deployment
  2. StatefulSet
  3. Daemon



### 基本命令

* 创建deployment

```bash
kubectl apply -f $file_path
```



* 查看所有deployment

```bash
kubectl get deployments
```



* 查看deployment

```bash
kubectl get deployment $deployment_name
```



* 查看deployment详情

```bash
kubectl describe deployment $deployment_name
```



* 删除deployment

```bash
kubectl delete deployment $deployment_name
```



### 高级命令

* 更新deployment

```bash
kubectl edit deployment.v1.apps/$deployment_name
```



* 查看上线状态

```bash
kubectl rollout status deployment/$deployment_name
```

