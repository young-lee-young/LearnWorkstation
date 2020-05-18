# 问题解决记录
--------------------------------------

## 创建管理员用户

```sh
python manage.py createsuperuser
# 输入用户名、邮箱、密码
```


## 命名空间的问题
namespace添加在include里面
path(r'app1', include('app1.urls', namespace='app1'))

>报错'Specifying a namespace in include() without providing an app_name'
>解决办法：在对应的app中urls.py文件urlpatterns前面加上app_name = 'app1'


>报错：AssertionError: `base_name` argument not specified, and could not automatically determine the name from the viewset, as it does not have a `.queryset` attribute.
>解决办法：在注册url时候加上base_name，
```
router = routers.DefaultRouter()
router.register(r'users', views.UserViewSet，base_name='users')
router.register(r'groups', views.GroupViewSet, base_name='groups')
```
>加上base_name后，访问加base_name的url会报错
>