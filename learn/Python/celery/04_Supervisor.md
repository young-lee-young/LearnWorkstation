# Supervisor使用


### 基础

* 安装

```sh
pip install supervisor
```

### 使用

* 将配置文件重新定位

```sh
echo_supervisord_conf > 重新定位的文件
```

* 启动supervisor

```sh
supervisord -c 配置文件位置

# 查看进程
ps -ef | grep supervisor
```

* 控制台操作

```sh
# 进入控制台
supervisorctl
```

* 关闭supervisor

在控制台中执行shutdown命令
