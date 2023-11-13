# Namespace


### 基础

* 作用：虚拟集群


### 基础命令

* 创建 namespace

```bash
kubectl create namespace $namespace_name

kubectl create -f $yaml_file
```


* 查看所有的 namespace

```bash
kubectl get namespaces
```


* 查看 namespace

```bash
kubectl get namespace $namespace_name
```


* 查看 namespace 详情

```bash
kubectl describe namespace $namespace_name
```


* 删除 namespace

```bash
kubectl delete namespace $namespace_name
```
