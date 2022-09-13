数据库（PostgreSQL, MySQL）

* PostgreSQL需要安装psycopg2,

```bash
pip install psycopg
```

* MySQL需要安装mysqlclient 

```sh
# 安装MySQL依赖
sudo apt-get install libmysqlclient-dev

pip install mysqlclient
```



在setting中配置数据库
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        # 数据库需要提前创建好
        'NAME': '数据库的名字',
        'USER': '数据库用户名',
        'PASSWORD': '数据库密码',
        'HOST': '127.0.0.1',
        'PORT': '3306',
    }
}




数据模型
models.py文件中创建一个类，继承models.Model，类就是一个数据表，在类里面创建字段
from django.db import models
class Liuyan(models.Model):
	# max_length是必须的
	username = models.CharField(max_length=32)
	fabiaotime = models.TimeField(auto_now=False, auto_now_add=True)
	content = models.TextField()

    def __str__(self):
    return self.username



1.7之前
python manage.py syncdb
1.7及以上创建数据模型（在数据库中生成数据表）
python manage.py makemigrations app名称   # 不加APP名字的话默认所有的APP都生成数据表，这个操作会在APP中migrations文件夹下创建一个0001_initial.py文件
python manage.py migrate
查看创建数据表的语句
python manage.py sqlmigrate app名字 0001


from myapp.models import Article
class Article(models.Model):
    title = CharField(max_lenth=30)
    see_num = IntegerField(default=0)

    def __str__(self):
        return self.title

    def create_article(self):
        Article.objects.create(title='xuwei', see_num=21)

    def get_article(self):
        Article.objects.get(name='xuwei')


新建对象的方法，前三种返回的是对象本身，最后一个返回的是一个元组
第一种
Article.objects.create(title='xuwei', see_num=21)
第二种
article = Article(title='xuwei', see_num=21)
article.save()
第三种
article = Article(title='xuwei')
article.see_num = 21
article.save()
第四种，这种相对较慢，返回一个元组，第一个参数是Article对象，第二个参数是True或者False，新建时返回True，已经存在返回False
Article.objects.get_or_create(title='xuwei', see_num=32):


获取对象的方法
Article.objects.all()
Article.objects.all()[:10]  # 切片操作，获取10个文章，不支持负索引
Article.objects.get(title='xuwei')

Article.objects.filter(title__exact='xuwei')    # 严格等于
Article.objects.filter(title__iexact='xuwei)    # 名称严格，不区分大小写
Article.objects.filter(title__contains='xuwei') # 包含xuwei
Article.objects.filter(title__icontains='xuwei')    # 名称中包含'xuwei'，且'xuwei'不区分大小写
Article.objects.filter(title__regex='^xuwei')   # 正则表达式查询
Article.objects.filter(title__iregex='^xuwei')   # 正则表达式不区分大小写
找到排除条件外的元素
Article.objects.exclude(title__contains='xuwei') # 排除xuwei之外的
Article.objects.filter(title__contains='xuwei').exclude(see_num=21)

删除一条数据
Article.objects.filter(title__contains='xuwei').delete()

更新数据
Article.objects.filter(title__contains='xuwei).update(title='markknopfler')


# QuerySet
items = Article.objects.all()得到的对象是可以迭代的
for item in items:
    print(item)

如果只是检查items中是否有对象
Article.objects.all().exists()
获取查询到的数量
Article..objects.count()
使用list(article)将QuerySet强制装换成列表


