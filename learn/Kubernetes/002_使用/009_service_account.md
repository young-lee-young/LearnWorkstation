# 用户、角色


### Service Account

* 查看所有的 service account

```bash
kubectl get serviceaccounts
```


* 创建 service account

```bash
kubectl create serviceaccount $serviceaccount_name
```


### Secret

* 创建 secret

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: postman-sa-secret
  annotations:
    kubernetes.io/service-account.name: forpostman
type: kubernetes.io/service-account-token
```

```bash
kubectl apply -f $yaml_file
```


* 查看所有的 secret

```bash
kubectl get secrets
```


* 查看 secret 详情

```bash
kubectl describe secret $secret_name
```


### ClusterRole

* 查看所有的 cluster role

```bash
kubectl get clusterroles
```


### RoleBinding

* 角色绑定

```bash
kubectl create rolebinding $rolebinding_name --clusterrole $cluster_role --serviceaccount $namespace:$service_account
```


* 查看所有的 role binding

```bash
kubectl get rolebindings
``` 


* 查看 role binding 详情

```bash
kubectl describe rolebinding $rolebinding_name
```
