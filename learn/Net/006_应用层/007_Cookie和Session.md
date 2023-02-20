# cookie 和 session

* 什么是会话

客户端打开与服务器的连接发出请求到服务器响应客户端请求的全过程


* 为什么需要 Cookie 和 Session

1. HTTP 为**无状态协议**，无法通过协议本身识别用户
2. 用户认识的需要，从而可以对用户的行为与权限作出判断，决定是执行或拒绝
3. 记录部分或全部用户信息，即用户行为追踪


* cookie 和 session

1. session 在服务器端，cookie 在客户端（浏览器）
2. session 可以放在 文件、数据库、或内存中都可以
3. session 的运行依赖 session ID，而 session ID 是存在 cookie 中的，也就是说，如果浏览器禁用了 cookie，同时 session 也会失效（但是可以通过其它方式实现，比如在 URL 中传递 session ID）
4. 用户验证这种场合一般会用 session

因此，维持一个会话的核心就是客户端的唯一标识，即 session id
