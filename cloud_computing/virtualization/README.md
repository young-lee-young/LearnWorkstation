# 虚拟化


### mac 生成

```python3
import hashlib
from datetime import datetime


def timestamp_ms():
    # 获取当前时间
    now = datetime.now()

    # 转换为毫秒级时间戳
    timestamp_ms = round(now.timestamp() * 1000)

    return str(timestamp_ms)

def md5(n):
    """MD5加密"""
    return hashlib.md5(n.encode()).hexdigest()


def gen_mac(m):
    mac = "52:54:00" + ":" + m[0:2] + ":" + m[2:4] + ":" + m[4:6]
    return mac

def main():
    n = timestamp_ms()
    m = md5(n)
    mac = gen_mac(m)
    print(mac)


if __name__ == "__main__":
    main()
```
