### Docker 文件系统

```bash
# 查看镜像详细信息
docker image inspect $image_id
# 可以看到 GraphDriver 有 Lower、Upper、Merge、Work 目录

# 到镜像目录下，overlay2 目录下一级
cd 85979fc2cfdc9be6b8537f3aa06804b6cf479beb6668d47f22420af890877e46
ls -alh
```


* diff 目录：存放当前层文件

查看镜像 GraphDriver 中 LowerDir 时候，看到都是上面所有层 diff 目录，使用 : 分割


* link 文件：当前层的的软连接

```bash
# 查看 diff 目录中文件
ls -alh diff

# 查看 link 文件内容
cat link
# 输出：S2AL7OU6U4QILWQC3VX5YETBSY

# 查看实际目录，可以看到是一个软链接，链接到 diff 目录
# 使用软连接目的：避免 mount 命令参数长度限制
ls -alh ../l/S2AL7OU6U4QILWQC3VX5YETBSY
# 输出：lrwxrwxrwx 1 root root 72 Mar 28 07:10 ../l/S2AL7OU6U4QILWQC3VX5YETBSY -> ../85979fc2cfdc9be6b8537f3aa06804b6cf479beb6668d47f22420af890877e46/diff
```


* lower 文件

```bash
# 查看 lower 文件，可以看到都是软连接
# 越靠后的越在底层，最后一个为最基础的镜像
cat lower
# 输出：l/Y727RC7XAM3DLFLFMTSCRS2LSU:l/MEJAWWODYIWOCSR6GLSKIPCTY3:l/C3DDPPGTDFADBO4RIRIGEA55RO:l/PVEKEJRU5U6YAQYCNBDK2RHWVT

ls -alh ../l/Y727RC7XAM3DLFLFMTSCRS2LSU
```


### Docker 元数据

目录：/var/lib/docker/image/overlay2

```bash
ls /var/lib/docker/image/overlay2
# 输出：distribution  imagedb  layerdb  repositories.json
```

* 镜像元数据

镜像元数据：imagedb/content/sha256

```bash
# 查看镜像元数据文件，文件为镜像的 ID，这个文件就是 docker image inspect $image_id 看到的信息
ls imagedb/content/sha256
```
