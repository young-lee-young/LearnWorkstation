# wsgi：web server gateway interface

描述Web Server（Gunicorn、uWSGI）如何与web框架交互，Web框架如何处理请求


### wsgi应用

* app.py文件

```python
def myapp(environ, start_response):
    print(environ)
    status = '200 OK'
    headers = [('Content-Type', 'text/html; charset=utf-8')]
    start_response(status, headers)
    return [b'<h1>Hello World</h1>']
```

* server.py文件

```python
from wsgiref.simple_server import make_server
from app import myapp


httpd = make_server('127.0.0.1', 8888, myapp)
httpd.serve_forever()
```
