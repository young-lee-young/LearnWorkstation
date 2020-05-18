# beego


### 安装

* beego安装

```sh
# 安装
go get github.com/astaxie/beego

# 升级
go get -u github.com/astaxie/beego
```

* bee工具安装

```sh
# 安装完成后会在GOPATH/bin生成bee文件
go get github.com/beego/bee

# 查看beego及bee版本
bee version
```


### 使用

* 创建项目

```sh
bee new 项目名称
```

* 运行项目

```sh
# 进入到项目目录
bee run
```

* 

### beego框架

* cache（缓存）

支持文件、内存、memcache、redis四种引擎

* config（配置文件）

支持ini、json、xml、yaml四种格式配置文件

* context

request、response

* httplibs

get、post、put、delete、head

* logs

多种输出引擎、异步输出

* orm

* session

* toolbox
