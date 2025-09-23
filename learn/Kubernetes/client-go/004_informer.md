# Informer

保证数据实时性、可靠性、顺序性


### 使用

* demo

```go
package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

func informerDemo(c *rest.Config) {
	clientSet, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	factory := informers.NewSharedInformerFactory(clientSet, time.Minute)

	// sharedIndexInformer
	podInformer := factory.Core().V1().Pods().Informer()

	_, err = podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			o := obj.(v1.Object)
			fmt.Printf("pod add to store: %s\n", o.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oo := oldObj.(v1.Object)
			no := newObj.(v1.Object)
			fmt.Printf("%s update to %s\n", oo.GetName(), no.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			o := obj.(v1.Object)
			fmt.Printf("pod delete from store: %s\n", o.GetName())
		},
	})

	factory.Start(stopCh)

	cache.WaitForCacheSync(stopCh, podInformer.HasSynced)
}
```


### 源码分析

```go
// tools/cache/shared_informer.go
type sharedIndexInformer struct {
    controller Controller
}

func (s *sharedIndexInformer) Run(stopCh <-chan struct{}) {
    fifo := NewDeltaFIFOWithOptions(DeltaFIFOOptions{
        KnownObjects:          s.indexer,
		EmitDeltaTypeReplaced: true,
        Transformer:           s.transform,
    })
}
```
