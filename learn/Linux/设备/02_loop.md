# loop 设备


### loop 设备作用

在文件和块设备之间建立映射关系


### 使用

* losetup

losetup 是管理 loop 设备的命令

* 示例

```bash
# 创建 disk.img 文件
dd if=/dev/zero of=/root/disk.img bs=1M count=1024

# 这个命令是将 loop0 设备和 disk.img 文件绑定
losetup /dev/loop0 /root/disk.img

# 将 disk.img 文件做成 ext4 文件系统
mkfs.ext4 /dev/loop0

# 挂载 loop 设备
mount /dev/loop0 /tmp/mnt/
```
