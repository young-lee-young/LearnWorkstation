# APIServer


### API Object

* API Object

Kubernetes 内部管理的基本元素，是 Kubernetes 在 ETCD 中信息存储单元

例如 Deployment、Pod、Service 都是 API Object


* API Group

一组 API Object 组成的一个具有共有性质的对象集合

例如，apps 这个 group，由 Deployment、ReplicaSet、StatefulSet 等 API Object 组成


* Legacy API Object 

Kubernetes 项目初始阶段所引入的 API Object，没有显示定义在 API Group 下面

例如 Pod、Event、Node 等
