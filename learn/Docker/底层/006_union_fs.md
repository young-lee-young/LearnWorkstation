# UnionFS


### 实验

有 A 目录，下面有文件 a、x，文件 a 内容为 a，文件 x 内容为 ax

有 B 目录，下面有文件 b、x，文件 b 内容为 b，文件 x 内容为 bx

* 实验

```bash
# 创建目录 C
mkdir C

# 将 A、B 目录联合挂载到 C 目录
mount -t aufs -o dirs=./A:./B none ./C
```
