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


### overlay 文件系统

* 实验

```bash
# 创建目录 & 进入目录
mkdir demo
cd demo

# 创建文件夹
mkdir A
mkdir B
mkdir C
mkdir work

# 写入文件
echo "from A" > A/a.txt
echo "from A" > A/b.txt
echo "from A" > A/c.txt
echo "from B" > B/a.txt
echo "from B" > B/d.txt
echo "from C" > C/b.txt
echo "from C" > C/e.txt

# 创建 merge 目录
mkdir /home/lee/merged

# 联合挂载
mount -t overlay overlay -o lowerdir=A:B,upperdir=C,workdir=work /home/lee/merged

# 查看
df -h
# Filesystem               Type      Size  Used Avail Use% Mounted on
# overlay                  overlay    19G  7.0G   11G  40% /home/lee/merged
```

* 分析

```bash
mount -t overlay overlay -o lowerdir=A:B,upperdir=C,workdir=work /home/lee/merged
```

文件 merge 的优先级

1. upper > lower
2. : 前面 > : 后面
