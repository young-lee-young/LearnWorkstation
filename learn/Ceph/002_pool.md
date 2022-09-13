# ceph pool


### 概念

pg: placement group

pgp: placement group for placement purpose


### pool命令

```bash
# 创建
ceph osd pool create $pool_name $pg_num $pgp_num replicated

# 启用
ceph osd pool application enable $pool_name rbd

# 列举
ceph osd pool ls

# 删除
ceph osd pool delete $pool_name $pool_name --yes-i-really-really-mean-it
```
