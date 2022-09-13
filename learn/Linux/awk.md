# awk命令使用


### 基本格式

```bash
awk 选项 '条件1 {动作1} 条件2 {动作2} ...' 文件名
```


### 处理流程

1. 读取第一行数据，并将第一行数据填入 $0、$1、$2变量中，其中 $0 代表一整行
2. 判断条件限制，是否需要进行后面的动作
3. 做完判断与动作
4. 执行下一行，继续1操作


### 选项

* -v：设置变量值


### 内置变量

1. NF：每一行拥有的列数
2. NR：目前所处理的行数
3. FS：目前的分隔字符，默认是空格


### 

```bash
awk -F '[ ,]' 'BEGIN {print "start process data -------"}{if ($1 < 2 || $4 == 23) {printf("filename: %s, lines: %s, columns: %s, name: %s %s\n", FILENAME, NR, NF, $2, $3)}} END {for (i = 0; i < NR; i++) print i}' awk_learn.txt
```
