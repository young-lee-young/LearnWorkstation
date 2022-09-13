# Django基本使用


### 虚拟环境

* 安装虚拟环境

```sh
pip install virtualenv
```

* 创建虚拟环境

```sh
virtualenv 虚拟环境名称

# 指定Python版本创建虚拟环境
virtualenv -p E:\Software\python279\python2.exe 虚拟环境名称
```

* 进入虚拟环境

```sh
# windows下进入虚拟环境
进入到文件夹中, 再进入到Scripts, 运行activate.bat文ls

# Mac、linux下进入虚拟环境
进入到bin目录下，source ./activate
```

* 退出虚拟环境

```sh
deactivate
```


### Django

* 安装Django

```sh
# 安装Django
pip install django
# 安装指定版本的Django
pip install django==1.8

# 查看当前所装的库
pip freeze

# 查看所有库
pip list

# 安装全部依赖库
pip install -r requirements.txt
```

* 创建与运行项目

```sh
# 新建一个项目
django-admin startproject 项目名

# 运行项目
python manage.py runserver 端口（默认8000端口）
python manage.py runserver 端口 --noreload --settings=resop.settings_test（不自动加载配置文件，指定配置文件）
```
    
* App

```sh
# 创建一个APP
python manage.py startapp app名称

# 将APP的名字加入到settings中
在APP目录下templates文件夹和static文件夹，分别存放HTML文件和静态文件
注意，Django默认是以settings中的APP顺序查找HTML文件，需要在每个APP下templates文件夹下再建一个以APP名字的文件夹，同时修改views中对应的路径

# 快速创建APP
django-admin startapp quickstart
```
