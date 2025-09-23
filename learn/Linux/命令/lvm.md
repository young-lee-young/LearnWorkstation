# LVM（logical volume manager）：逻辑卷管理


### vg：volume group

```bash
# 创建物理卷组，并将物理卷加如卷组
# vgcreate 物理卷组名 物理磁盘设备1 物理磁盘设备2
vgcreate lee /dev/nvme0n2 /dev/nvme0n3 # 将物理设备加入到一个 vg

# 查看物理卷组
vgs
vgscan
vgdisplay

# 将物理卷加入到物理卷组
# vgextend 物理卷组名 物理卷设备
vgextend lee /dev/nvme0n4

# 将物理卷从物理卷组移除
# vgreduce 物理卷组名 物理卷设备
vgreduce lee /dev/nvme0n4

# 删除物理卷组
# vgremove 物理卷组名
vgremove lee
```


### pv：physical volume

```bash
# 创建物理卷
# pvcreate 物理设备
pvcreate /dev/nvme0n2

# 查看物理卷
pvs
pvscan
pvdispaly

# 删除物理卷
# pvremove 物理设备
pvremove /dev/nvme0n2
```


### lv：logical volume

```bash
# 创建逻辑卷
# lvcreate -n 逻辑卷名 -L 逻辑卷大小 物理卷组名
lvcreate -n young -L 10G lee
lvcreate -n young -l 100%free lee
# 创建后会生成 /dev/物理卷组名/逻辑卷名 的逻辑设备（如上面的命令会生成 /dev/lee/young 的逻辑设备）

# 查看逻辑卷
lvs
lvscan
lvdisplay

# 扩容逻辑卷
# lvextend -L +扩容大小 逻辑卷设备
lvextend -L +10G /dev/lee/young

# 调整逻辑卷大小
# lvresize -L 调整目标大小 逻辑卷设备
lvresize -L 10G /dev/lee/young

# 删除逻辑卷
# lvremove 逻辑卷设备
lvremove /dev/lee/young
```


### 使用逻辑卷

```bash
# 做文件系统
# mkfs.文件系统类型 逻辑卷设备
mkfs.xfs /dev/lee/young

# 挂载
# mount 逻辑卷设备 目的路经
# 如 挂载 docker 的目录
mount /dev/lee/young /var/lib/docker

# 查看文件系统信息
df -Th
```


# thin LVM：精简逻辑卷管理

* 创建 vg

* 在物理卷组中创建精简池

```bash
# lvcreate --thinpool -n 精简池名 -L 精简池大小 物理卷组名
lvcreate --thinpool pool -L 10G lee
```

* 查看精简池

```bash
lvdisplay
```

* 创建逻辑卷

```bash
# lvcreate --thin -V 逻辑卷大小 -n 逻辑卷名 物理卷组名/精简池名
# 会创建 /dev/物理卷组名/逻辑卷名 设备
lvcreate --thin -V 5G -n pool_lv lee/pool
```

* 扩容精简池

```bash
# lvextend -L +扩容大小 精简池设备
lvextend -L +5G /dev/lee/pool
```

* 删除精简池

```bash
# lvremove 精简池设备
lvremove /dev/lee/pool
```
