# rpm命令使用（用于安装.rpm文件）


### 基本格式

```bash
rpm 选项 文件名
```


### 选项

* -i：install（安装）
* -U：update（安装或更新）
* -e：erase（卸载）
* -q：query(查询)
* -V：verify（校验）

* -v：显示详细信息
* -h：hash marks：输出进度条


### 安装

```bash
rpm -ivh 包名
```


### 升级

```bash
# 升级或安装
rpm -Uvh 包名

# 仅升级，若没有安装，不升级
rpm -Fvh 包名
```


### 查询

```bash
rpm -qa
```


### 卸载

```bash
rpm -e 包名
```
