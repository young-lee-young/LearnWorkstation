# minikube 搭建集群


### 安装 minikube

地址：https://minikube.sigs.k8s.io/docs/


### 初始化集群

```bash
minikube start --kubernetes-version=v1.30.0 --driver=docker --cni=calico --force
```
