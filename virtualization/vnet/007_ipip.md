

### 

```bash
modprobe ipip
lsmod | grep ipip

# 添加tun设备，设置模式为ipip，设置隧道两端，隧道外层IP
ip tunnel add tun1 mode ipip remote 10.10.20.2 local 10.10.10.2
# 设置隧道内层IP
ip addr add 10.10.100.10 peer 10.10.200.10 dev tun1
```
