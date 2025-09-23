# 四种客户端


### RESTClient

对 HTTP Request 的封装，实现 RESTful 风格的 API，其他客户端基于 RESTClient 实现

* 实现

在 rest 目录下

* demo

```go
package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

func restClientDemo(c *rest.Config) {
	c.APIPath = "api"
	c.GroupVersion = &corev1.SchemeGroupVersion
	c.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(c)
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}
	err = restClient.Get().Namespace("").Resource("pods").VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).Do(context.Background()).Into(result)
	if err != nil {
		panic(err)
	}

	for _, item := range result.Items {
		fmt.Printf("NAMESPACE: %s, NAME: %s, STATUS: %s\n", item.Namespace, item.Name, item.Status.Phase)
	}
}
```

### ClientSet

多种 Resource 客户端的集合，不能直接访问 CRD 资源，ClientSet 需要生成代码才能操作 CRD 资源

* 实现

在 kubernetes 目录下

* demo

```go
package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func clientSetDemo(c *rest.Config) {
	clientSet, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err)
	}

	podClient := clientSet.CoreV1().Pods(corev1.NamespaceAll)
	result, err := podClient.List(context.Background(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	for _, item := range result.Items {
		fmt.Printf("NAMESPACE: %s, NAME: %s, STATUS: %s\n", item.Namespace, item.Name, item.Status.Phase)
	}
}
```

### DynamicClient

处理 K8S 内置资源和 CRD 资源

* 实现

在 dynamic 目录下

* demo

```go
package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func dynamicClientDemo(c *rest.Config) {
	dynamicClient, err := dynamic.NewForConfig(c)
	if err != nil {
		panic(err)
	}

	resource := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}

	unstructObj, err := dynamicClient.Resource(resource).Namespace(corev1.NamespaceAll).List(context.Background(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}

	for _, item := range podList.Items {
		fmt.Printf("NAMESPACE: %s, NAME: %s, STATUS: %s\n", item.Namespace, item.Name, item.Status.Phase)
	}
}
```

### DiscoveryClient

发现 kube-apiserver 所支持的资源组（Group）、资源版本（Version）、资源信息（Resource）

会缓存在本地 ~/.kube/cache 和 ~/.kube/http-cache

* 实现

在 discovery 目录下

* demo

```go
package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

func discoveryDemo(c *rest.Config) {
	client, err := discovery.NewDiscoveryClientForConfig(c)
	if err != nil {
		panic(err)
	}

	_, apiResourceList, err := client.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	for _, item := range apiResourceList {
		gv, err := schema.ParseGroupVersion(item.GroupVersion)
		if err != nil {
			panic(err)
		}

		fmt.Printf("GROUP %s, VERSION: %s\n", gv.Group, gv.Version)
		for _, resource := range item.APIResources {
			fmt.Printf("NAME: %v\n", resource.Name)
		}
		fmt.Println("----------------------------------------------")
	}
}
```
