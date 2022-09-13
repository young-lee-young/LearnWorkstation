### 实现

mysqldump 全量备份后，使用 BinLog 作为增量


### 备份

* 全量备份

```bash
# --flush-logs：备份后切换 BinLog 文件
# --master-data=2：记录切换后的 BinLog 文件名
mysqldump -u$USER -p$PASSWORD --databases $DATABASE --single-transaction --flush-logs --master-data=2 > $FILENAME.sql
```

* 增量备份

```bash
# 切换 BinLog 文件
mysqladmin -u$USER -p$PASSWORD flush-logs

# 备份 BinLog
```


### 还原

* 恢复全量备份

```mysql
mysql> source $FILENAME.sql
```

* 还原 BinLog 增量

```bash
# 可以有多少 BinLog，用空格隔开
mysqlbinlog $BinLog_FILENAME $BinLog_FILENAME | mysql -u$USER -p$PASSWORD
```
