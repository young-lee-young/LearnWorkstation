### mysqldump

输出的备份内容为 SQL 语句，平衡了阅读和还原；SQL 语句占用小

* 原理

```mysql
# SQL_NO_CACHE 查询出的数据不会进入 SQL 缓存
SELECT SQL_NO_CACHE * FROM $DATABASE;
```


### 备份还原

* 备份

```bash
# --single-transaction，在 RR 级别下进行
mysqldump -u$USER -p$PASSWORD --databases $DATABASE --single-transaction > $FILENAME.sql
```

* 还原

```bash
mysql> source $FILENAME.sql
```


### 注意

* 在 MyISAM 下备份

--lock-all-tables：使用 FTWRL 锁所有表

--lock-tables：使用 READ LOCAL 锁当前库的表


### 缺点

1. 导出逻辑数据，备份较慢
2. 还原执行 SQL 语句，速度较慢
