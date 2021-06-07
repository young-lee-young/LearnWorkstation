
* 创建超级用户

```sh
python manage.py createsuperuser
# 输入用户名，输入邮箱，输入密码，再次输入密码
```



* 进入admin后台

```sh
域名/admin，输入用户名和密码
```



管理后台，在APP下有一个admin.py，修改文件
from django.contrib import admin
from . import models

admin.site.register(models.Liuyan)
