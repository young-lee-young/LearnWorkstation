# bximage 使用


### 使用

* 创建硬盘镜像

```bash
# 创建硬盘镜像：硬盘大小16M，扇区大小215字节
bximage -q -mode=create -hd=16 -sectsize=512 -imgmode=flat master.img
```


* 将 boot.bin 写入硬盘镜像

```bash
dd if=boot.bin of=master.img bs=512 count=1 conv=notrunc
```
