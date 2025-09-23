### 配置

* demo

```go
package demo

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func generateConfig() *rest.Config {
	masterURL := ""
	path := "/Users/lee/.kube/config"
	c, err := clientcmd.BuildConfigFromFlags(masterURL, path)
	if err != nil {
		panic(err)
	}
	return c
}
```
