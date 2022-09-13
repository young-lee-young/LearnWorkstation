# yum（Yellow dog Updater Modified）命令使用


### 基本格式

```bash
yum 选项 命令 操作对象
```


### 选项

* -h：帮助
* -q：不显示安装过程
* -y：安装过程中提示全部选择yes


### 安装包

```bash
yum search 包名

yum install -y 包名
```


### 查看包

```bash
yum list

yum info 包名
```


### 更新包

```bash
yum check-update

yum update 包名
```


### 删除包

```bash
yum deplist 包名

yum remove 包名
```


### 组操作

```bash
yum grouplist

yum groupinfo 组名

yum groupinstall 组名

yum groupupdate 组名

yum groupremove 组名
```


### 缓存操作

```bash
# 清除缓存目录(/var/cache/yum)下的软件包
yum clean packages

# 清除缓存目录下的headers
yum clean headers

# 清除缓存目录下旧的headers
yum clean oldheaders

# yum clean packages + yum clean oldheaders
yum clean
```


### 修改yum源

* 备份/etc/yum.repos.d/CentOS-Base.repo
* 下载国内源的repo文件，放入/etc/yum/repos.d目录下
* 执行以下命令，生成缓存

```bash
yum clean all

yum makecache
```
