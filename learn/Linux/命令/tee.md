# tee 命令


### 使用说明

从标准输入读取数据，写入文件的同时把数据输出到标准输出

```bash
# echo hello 会把 hello 输出到标准输出
# | 管道符会把 echo 命令的标准输出连接到下一个命令的标准输入
# tee 会读取标准输入，并且把标准输入 hello 写入到 hello.txt，并且输出到标准输出，所以会在控制台打印出 hello
echo hello | tee hello.txt

# 输出：hello
```

```bash
# -a 参数：以追加的形式写入文件
echo hello | tee -a hello.txt
```
