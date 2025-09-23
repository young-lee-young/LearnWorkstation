# controller


### 开发

* 移动文件（业界常规做法）

将 internal/controller/$kind_controller.go 和 internal/controller/$kind_controller_test.go 移动到目录 pkg/$kind/controller

```bash
mkdir -p pkg/book/controller
mv internal/controller/book_controller.go pkg/book/controller
mv internal/controller/book_controller_test.go pkg/book/controller
```


* 运行 Controller

```bash
make run
```


### 部署

* 打包镜像

```bash
# make docker-build docker-push IMG=$registry/$project:$tag
make docker-build docker-push IMG=hello/demo:v0.1
```


* 部署

```bash
# make deploy IMG=$registry/$project:$tag
make deploy IMG=hello/demo:v0.1

# 输出
namespace/helloworld-system created
customresourcedefinition.apiextensions.k8s.io/books.lee.com.lee unchanged
serviceaccount/helloworld-controller-manager created
role.rbac.authorization.k8s.io/helloworld-leader-election-role created
clusterrole.rbac.authorization.k8s.io/helloworld-book-editor-role created
clusterrole.rbac.authorization.k8s.io/helloworld-book-viewer-role created
clusterrole.rbac.authorization.k8s.io/helloworld-manager-role created
rolebinding.rbac.authorization.k8s.io/helloworld-leader-election-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/helloworld-manager-rolebinding created
deployment.apps/helloworld-controller-manager created
```


* 取消部署

```bash
make undeploy
```
