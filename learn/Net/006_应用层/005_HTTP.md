# HTTP（Hyper Text Transfer Protocol，超文本传输协议）


### 基础

* 什么是超文本

超越了普通文本的文本，是文字、图片、音频和视频等的混合体，最关键的是含有"超链接"，能够从一个超文本跳跃到另一个超文本，形成复杂的网状的结构关系


### HTTP 协议特点

* HTTP 使用的是可靠的传输协议

传输层使用 TCP 协议


* 简单快速

客户端向服务器请求时，只需传送请求方法和路径

常用方法：GET、POST、HEAD、DELETE、PUT


* 灵活

HTTP 允许传输任意类型的数据对象


* 无状态

HTTP 协议自身不对请求和响应之间的通信状态进行保存


### 输入网址后发生了什么

1. 浏览器解析域名，解析先后顺序为：浏览器缓存 > hosts > DNS 服务器
2. TCP 三次握手，建立连接
3. 发送 HTTP 请求
4. 服务端响应数据
5. 浏览器收到数据后渲染页面
6. TCP 四次挥手，断开连接


### HTTP 报文

* 请求报文

请求行：请求方法、URL、HTTP协议及版本
请求头：key：value 回车符 换行符（\r\n）
       可以有很多组 key-value 
回车符 换行符（\r\n）
请求体


* 响应状态码

200：OK
301：永久重定向
302：暂时重定向
400：Bad Reqeust
401：Unauthorized
403：Forbidden
404：Not Found
405：Method Not Allowed
500：Internal Server Error
501：Not Implemented
502：Bad Gateway


### 问题

* 如何区分请求头和请求体
