# -*- coding: utf-8 -*-
# @Time:    2019/9/23 下午2:01
# @Author:  leeyoung
# @File:    myapp.py
# @Content:


def application(environ, start_response):
    # print(environ)
    status = '200 OK'
    headers = [('Content-Type', 'text/html; charset=utf-8')]
    start_response(status, headers)
    return [b'<h1>old driver</h1>']
