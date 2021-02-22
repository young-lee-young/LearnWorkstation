
### 添加一台机器（首先两台机器要能ping通）

1. 编辑/etc/ansible/hosts
2. 添加本机的公钥到目标机器的authorized_keys
3. 添加本机私钥到ansible
4. ansible all -m ping			# 测试连接


### ansible
