# views视图操作

### RequestContext使用
如果在很多视图中都是需要相同的数据，未防止数据冗余，使用RequestContext
```
from django.template import RequestContext

# 定义公共数据，context处理器
def public_data():
    data = {
        'name': 'lee',
        'address': 'beijing',
    }
    return data


# view视图
def commonview(request):
    request_context = RequestContext(request, {'age': 22}, processor=[public_data])
    # RequestContext接收3个参数, 第一个是request, 第二个是数据，第三个是context处理器列表
    return render_to_response('main.html', Request)
    # render_to_response参数是模板名, 数据, 其他 
    # render的参数是request, 模板名, 数据, 其他
```
context_instance(RequestContext(request))



from django.shortcuts import render
from . import models
def callme(request):
	liuyan = models.Liuyan.objects.get(pk=1)
	return render(request, 'callme/callme.html', {'liuyan': liuyan})


第一步导入HttpResponse
from django.http import HttpResponse

def return_func():
    html = '<html><body>hello %s, welcome</body></html>' % leeyoung
    retrun HttpResponse(html)
