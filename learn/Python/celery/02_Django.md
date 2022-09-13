# Django中使用celery


### 包安装

```sh
pip install django-celery
```

### 步骤

1. 在settings.py文件中添加djcelery APP，并配置BROKER等
2. 在settings.py同级目录下创建celeryconfig.py文件，并配置该文件
3. 


# 命令

* 启动worker

```sh
python manage.py celery worker -l INFO
```

* 启动beat

```sh
python manage.py celery beat -l INFO
```
