# HTTP（Hyper Text Transfer Protocol，超文本传输协议）


### 超文本

超越了普通文本的文本，是文字、图片、音频和视频等的混合体，最关键的是含有"超链接"，能够从一个超文本跳跃到另一个超文本，形成复杂的网状的结构关系


### Cookie和Session

Cookie会保存客户端信息，包括session
Session 保存在客户端，


* 输入网址后发生了什么

1. 浏览器解析域名
2. DNS查询IP地址（浏览器缓存，hosts，DNS服务器）
3. 建立连接HTTP请求
4. 反向代理Nginx
5. uwsgi、gunicorn
6. web app响应
7. TCP挥手
