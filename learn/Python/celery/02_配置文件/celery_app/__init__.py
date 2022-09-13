# -*- coding: utf-8 -*-
# @Time    : 2018-10-08 20:02
# @Author  : yao.li
# @Content :

from celery import Celery

app = Celery('demo')
# 通过Celery实例加载配置模块
app.config_from_object('celery_app.celeryconfig')
