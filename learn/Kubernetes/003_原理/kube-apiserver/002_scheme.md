# Scheme


### 代码位置

/kubernetes/staging/src/k8s.io/apimachinery/pkg/runtime


### 概念

* Version

每个 API Group 都有多个版本，每个 Version 包含多个 kind（一个 kind 可能出现在多个 version 下）


* GVK

Group、Version、Kind 确定唯一一个 Kind


* Type

API Object 结构体类型


### Scheme 作用

* GVK 和 Type 之间的转换

Scheme 内部有两个 map，分别 gvk -> type，type -> gvk，两者可以互相找到


* API Object 默认值

1. API Object 使用者在操作 object 时，不太可能给出所有属性值
2. Object 从一个 Version 转换到另一个 Version 时，可能需要为不存在对应关系的字段填值


* 内外部 Version 之间的转换

所有的外部 Version 都会被转换为内部 Version
