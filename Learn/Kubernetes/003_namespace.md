# namespace



### 基础

* 作用

  虚拟集群



### 基本命令

* 创建namespace

```bash
kubectl create namespace test-lee

kubectl create -f $file_path
```



* 查看所有namespace

```bash
kubectl get namespaces
```



* 查看namespace

```bash
kubectl get namespace $namespace_name
```



* 查看namespace详情

```bash
kubectl describe namespace $namespace_name
```



* 删除namespace

```bash
kubectl delete namespace $namespace_name
```

