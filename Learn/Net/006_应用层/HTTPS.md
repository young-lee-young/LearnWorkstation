# HTTPS（Hyper Text Transfer Protocol over Secure Socker Layer，安全超文本传输协议）


### SSL（Secure Sockets Layer，安全套接层）

对称加密：加密方和解密方使用同一个密钥
非对称加密：公钥和私钥存在关系，公钥进行传输，对方用公钥加密，用私钥加密hash值，将加密后的密文和加密的hash值传输过来，本方使用私钥解密，本地运算一个hash值，之后用对方的公钥解密hash值
