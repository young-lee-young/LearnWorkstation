### 物理备份

1. 直接备份 InnoDB 底层数据文件
2. 导出不需要转换，速度快
3. 工作时对数据库的压力较小
4. 更容易实现增量备份


### 备份

* 物理 + 热 + 全量 备份

1. 启动 RedoLog 监听线程，开始收集 RedoLog
2. 拷贝 ibd 数据文件
3. 停止收集 RedoLog
4. 加 FTWRL 锁拷贝 frm 元数据文件

* 物理 + 热 + 增量 备份

确定增量：根据每个页的 LSN 号，确定页的变化


### 还原

还原 ibd 文件，重放 RedoLog


### 使用

* 全量备份

```bash
innobackupex --user=$USER --password=$PASSWORD $BACKUP_DIR
```

* 增量备份

```bash
innobackupex --user=$USER --password=$PASSWORD --incremental $BACKUP_DIR --incremental-basedir='$BACKUP_DIR/XXXX-XX-XX/'
```

* 增量备份合并至全量备份

```mysql
innobackupex --apply-log $BACKUP_DIR/XXXX-XX-XX/ --incremental-dir=$BACKUP_DIR/YYYY-YY-YY/
```

* 还原（停掉 mysqld）

```bash
innobackupex --copy-back $BACKUP_DIR/XXXX-XX-XX/
```
