
# Linux服务器部署
-----------

### 环境安装
安装Python3相关依赖库：sudo apt-get install python3-dev build-essential libssl-dev libffi-dev libxml2 libxml2-dev libxslt1-dev zlib1g-dev
安装Python3：sudo apt-get install python3
安装pip：sudo apt-get install python3-pip


重启过程中可能会遇到unable to resolve host的问题，解决方法是修改/etc/hosts文件，在127.0.0.1 localhost下面加一句127.0.1.1 hostname，hostname是主机名，查看文件/etc/hostname获得
改表法：进入数据库，use mysql; update user set host = '%' where user = 'root'; flush privileges;



### 解析DNS域名

nginx是一个轻量级的web服务器，安装方式sudo apt-get install nginx
启动服务sudo /etc/init.d/nginx start
停止服务sudo /etc/init.d/nginx stop
重启服务sudo /etc/init.d/nginx restart


安装uwsgi，pip install uwsgi
开启uwsgi，sudo uwsgi --ini uwsgi.ini
关闭uwsgi，sudo killall -9 uwsgi
