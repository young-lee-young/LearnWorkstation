# Service


### 基础命令

* 查看所有 service

```bash
kubectl get services
```


* 查看 service

```bash
kubectl get service $service_name
```


* 查看 service 详情

```bash
kubectl describe service $service_name
```


* 删除 service

```bash
kubectl delete service $service_name
```


* 查看 service 代理的 pod

````bash
kubectl get endpoints $service_name
````
