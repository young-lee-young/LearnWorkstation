# Webhook


### 基础

* 作用

修改和验证


* 流程

user -> api server -> webhook -> controller


### 开发

* 准备工作

```bash
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.
2.0/cert-manager.yaml
```


* 创建 webhook

```bash
kubebuilder create webhook --group lee --version v1beta1 --kind Book --defaulting --programmatic-validation
```
