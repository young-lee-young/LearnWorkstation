### 基础命令

```bash
# 查看kubectl版本
kubectl version

# 查看集群信息
kubectl cluster-info

# 查看 API 版本
kubectl api-versions
```


* 查看日志：-v=0-9

```bash
kubectl get pods -v=9
```


* 监听模式：-w

```bash
kubectl get pos $pod_name -w
```
