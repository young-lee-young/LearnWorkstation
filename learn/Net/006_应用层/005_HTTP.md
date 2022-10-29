# HTTP（Hyper Text Transfer Protocol，超文本传输协议）


### 基础

* 什么是超文本

超越了普通文本的文本，是文字、图片、音频和视频等的混合体，最关键的是含有"超链接"，能够从一个超文本跳跃到另一个超文本，形成复杂的网状的结构关系


### HTTP协议特点

* HTTP 使用的是可靠的传输协议

传输层使用 TCP 协议


* 简单快速

客户端向服务器请求时，只需传送请求方法和路径

常用方法：GET、POST、HEAD、DELETE、PUT


* 灵活

HTTP 允许传输任意类型的数据对象


* 无状态

HTTP 协议自身不对请求和响应之间的通信状态进行保存


### cookie 和 session

1. session 在服务器端，cookie 在客户端（浏览器）
2. session 可以放在 文件、数据库、或内存中都可以
3. session 的运行依赖 session ID，而 session ID 是存在 cookie 中的，也就是说，如果浏览器禁用了 cookie，同时 session 也会失效（但是可以通过其它方式实现，比如在 URL 中传递 session ID）
4. 用户验证这种场合一般会用 session

因此，维持一个会话的核心就是客户端的唯一标识，即 session id


### 输入网址后发生了什么

1. 浏览器解析域名，解析先后顺序为：浏览器缓存 > hosts > DNS服务器
2. TCP 三次握手，建立连接
3. 发送 HTTP 请求
4. 服务端响应数据
5. 浏览器收到数据后渲染页面
6. TCP 四次挥手，断开连接
