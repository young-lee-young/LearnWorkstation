# Kubernetes 基础


### 组件

* master（控制节点）

APIServer：API 请求服务

kube-scheduler：调度

kube-controller-manager：容器编排（编排：处理容器之间的各种关系）

整个集群的持久化数据，保存在 etcd 中


* node（计算节点）

kubelet：和容器运行时交互，通过 CRI（container runtime interface）远程过程调用

```bash
sudo systemctl status kubelet
```

kube-proxy：pod 网络


### 基础命令

```bash
# 查看kubectl版本
kubectl version

# 查看集群信息
kubectl cluster-info

# 查看 API 版本
kubectl api-versions
```


### 配置 dashboard

```bash
# 查看当前运行上下文
kubectl config current-context

# 切换 kubernetes 运行上下文至 docker-for-desktop
kubectl config use-context docker-for-desktop

# 查看集群状态
kubectl cluster-info
kubectl get nodes

# 部署 kubernetes dashboard
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
# 或
kubectl create -f kubernetes-dashboard.yaml

# 查看dashboard pod 
kubectl get pod --namespace=kube-system | grep dashboard

# 将本地端口转发到pod，浏览器访问8443端口
kubectl port-forward 上一步pod-namespace 8443:8443 --namespace=kube-system

# 获取令牌
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | awk '/^deployment-controller-token-/{print $1}') | awk '$1=="token:"{print $2}'
```
