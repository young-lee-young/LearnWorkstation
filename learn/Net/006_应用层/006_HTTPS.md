# HTTPS（Hyper Text Transfer Protocol over Secure Socker Layer，安全超文本传输协议）


### SSL vs TSL

* SSL（Secure Sockets Layer，安全套接字层）

SSL 位于 7 层模型的会话层


* TLS（Transport Layer Security，安全传输层协议）


* TLS 和 SSL 的关系

TLS is actually just a more recent version of SSL. It fixes some security vulnerabilities in the earlier SSL protocols

翻译：TLS 只是 SSL 的最新版本


* 发展历史

SSL 1.0 – never publicly released due to security issues.
SSL 2.0 – released in 1995. Deprecated in 2011. Has known security issues.
SSL 3.0 – released in 1996. Deprecated in 2015. Has known security issues.
TLS 1.0 – released in 1999 as an upgrade to SSL 3.0. Planned deprecation in 2020.
TLS 1.1 – released in 2006. Planned deprecation in 2020.
TLS 1.2 – released in 2008.
TLS 1.3 – released in 2018.


* SSL/TLS 作用

1. 加密：
2. 鉴权：ensure that data is being sent to and received from the correct server, rather than a malicious “man in the middle.”（确保数据是从正确的服务端接收和发送，防止中间人假冒）
3. 数据完整性：ensure that there’s no loss or alteration of data during transport by including a message authentication code, or MAC.（确保数据没有丢失或者被篡改）


### 加密

* 对称加密

加密、解密使用同一个密钥


* 对称加密的问题

1. 密钥传输过程中可能被窃取


* 非对称加密

加密、解密使用不同的密钥，一把作为公开的公钥，一把作为保密的私钥，公钥加密的信息，只有私钥可以解开


* 非对称加密的问题

1. 非对称加密算法复杂，速度慢


* 混合加密

使用非对称加密交换密钥，之后使用对称加密交换数据


### 数字证书

* 证书来源

由权威的、受信任的证书颁发机构（CA）授予的


* 证书作用

1. 服务器向浏览器证明自己的身份
2. 把公钥传给浏览器


### 握手

* TLS 握手

TLS 有四次握手，四次握手在 TCP 三次握手之后进行


### 问题

* 为什么看到的都是 SSL 证书，不是 TLS 证书

[Why Aren't We All Using TLS Certificates Then 为什么我们不使用 TLS 证书](TLS证书.png)

翻译：使用什么样的协议取决于你的服务，虽然你看到的是 SSL 证书，但实际上使用的可能是 TLS 协议
