# 监控工具Flower使用


### 基础

* 介绍

Flower是使用Tornado写的

* 安装

```sh
pip install flower
```

### 使用

* 启动

```sh
# 常规启动
celery flower --port=5555 --broker=redis://username:password@localhost:6379/1
celery flower --port=5555 --broker=amqp://leeyoung:leeyoung@60.205.177.93:5672/hello/

# Django启动
python manage.py celery flower

# 浏览器访问localhost:5555
```
