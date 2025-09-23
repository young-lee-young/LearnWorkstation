
```bash
# 启动 http server
kubectl proxy


# 获取 pod
curl -v http://127.0.0.1:8001/api/v1/namespaces/kube-system/pods


# watch 模式
curl -v http://127.0.0.1:8001/api/v1/namespaces/kube-system/pods?watch=true
```
