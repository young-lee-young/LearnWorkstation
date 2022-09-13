# -*- coding: utf-8 -*-
# @Time:    2019/9/23 下午2:02
# @Author:  leeyoung
# @File:    server.py
# @Content:

from wsgiref.simple_server import make_server
from app import application


httpd = make_server('127.0.0.1', 8888, application)
httpd.serve_forever()
