# labels

### 查看pod

```bash
# 显示labels
kubectl get pods --show-labels

# 显示某个标签
kubectl get pods -L env
```

### 添加修改标签

```bash
# 添加新标签
kubectl label pod $pod_name 标签名=标签值

# 修改标签
kubectl label pod $pod_name 标签名=标签值 --overwirte
```


### 使用标签过滤pod

```bash
# 包含 env 标签的pod
kubectl get pods -l env

# 不包含 env 标签的pod
kubectl get pods -l '!env'

# env = dev的pod
kubectl get pods -l env=dev

# env != dev的pod
kubectl get pods -l env!=dev

# env in (dev, pro)
kubectl get pods -l 'env in (dev,pro)'

# env notin (dev, pro)
kubectl get pods -l 'env notin (dev, pro)'

# 多标签
kubectl get pods -l env=dev,creation=manual
```
