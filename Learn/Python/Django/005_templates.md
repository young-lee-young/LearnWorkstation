
django的模板语言

注释
{#单行注释#}
{%comment%}
    多行注释
{%endcomment%}

变量
{{ person }}
{{ person.name }}

标签
{% 关键字引导的程序逻辑 %}
标签中的关键字包括：for, endfor, block, endblock, if, elif, else, endif, in, trans, as, with, extends等
{% for line in datas %}
<tr>
    <td>{{ line.title}}</td>
    <td align="center">{{ line.time }}</td>
    <td>{{ line.content }}</td>
</tr>
{% endfor %}

过滤器
{% 变量|过滤标签 %}
对变量的值进行修饰
lower, escape, linebreaksm, date, length
{{ person.name|lower }}
{{ my_date|date:'Y-m-d' }}




包含通用文件
<!DOCTYPE html>
<html>
<head>
    <title>{% block title %}标题{% endblock %}</title>
</head>
<body>
    # 包含的文件
    {% include 'nav.html' %}
    {% block content %}
    <div>内容</div>
    {% endblock %}
    {% include 'bottom.html' %}
</body>
</html>


继承文件
{% extends 'base.html' %}
{% block title %}扩展标题{% endblock %}
{% block content %}
{% include 'add.html' %}
这里是首页
{% endblock %}



view中的代码
def home(request):
    info_dict = {'name': 'lee', 'address': 'beijing'}
    return render(request, 'home.html', {'info_dict': info_dict})

home.html中代码
姓名：{{ info_dict.name }} 地址：{{ info_dict.address }}

 遍历字典
{% for key, value in info_dict.items %}
    {{ key }}: {{value}}
{% endfor %}


遍历列表
def home(request):
   num_list = map(str, ranage(100))
   return render(request, 'home.html', {'List': List})

{% for item in List %}
    {{ item }},
{% end for }

判断是否便利到最后一个元素
{% for item in List %}
    {{ item }}{% if not forloop.last %},{% endif %} # 不是最后一个就加一个,
{% endfor %}

常用在for循环中的语句
forloop.counter # 索引从1开始
forloop.counter1 # 索引从0开始
forloop.revcounter  # 索引从最大长度到1
forloop.revcounter0 # 索引从最大长度到0
forloop.first   # 当遍历到第一项是为真
forloop.last    # 当遍历到最后一项为真
forloop.parentloop  # 用在嵌套的for循环中，获取上一层for循环的forloop


当循环的列表可能为空时
{ul}
{% for info in info_list %}
    <li>{{ athlete.name }}</li>
{% empty %}
    <li>列表为空</li>
{% endfor %}
</ul>


url取别名
{% url 'home' arg1 arg2 as new_url %}

html中
<a href='{ new_url }' >链接{{ new_url}}</a>


逻辑操作
{% if score >= 90%}
    优秀
{% elif score >=80 %}
    良好
{% elif score >= 0 and score < 80 %}
    及格
{% else %}
不及格end
{% endif %}

判断一个元素是不是在列表里
<% if 'lee' in List %>
'lee'在列表里
{% endif %}


获取网址和当前用户
{{ request.user }}
{% if request.user.is_authenticated %}
    {{ request.user.username }}, 您好
{% else %}
    请登录
{% endif %}


获取当前的网址
{{ request.path }}

获取当前的get参数
{{ request.GET.urlencode }}
