### OUTFILE

MySQL 原生的 SQL 指令，最原始的逻辑备份方式，备份的功能和效果取决于如何写 SQL


### 命令

* 查看导出路径

```mysql
mysql> SHOW VARIABLES LIKE '%secure%';
+--------------------------+-------+
| Variable_name            | Value |
+--------------------------+-------+
| secure_file_priv         |       |
+--------------------------+-------+
```

* 将查询结果导出文件

```mysql
BEGIN;

SELECT * INTO OUTFILE '/opt/homebrew/var/mysql/t3_1' FROM t3;

# 设置分隔符
SELECT * INTO OUTFILE '/opt/homebrew/var/mysql/t3_2' fields terminated by ',' FROM t3;

# 设置换行符
SELECT * INTO OUTFILE '/opt/homebrew/var/mysql/t3_2' lines terminated by ',' FROM t3;

COMMIT;
```


### 缺陷

1. 输出文本比较简略
2. 很难进行还原，现在往往用来简单导出数据
