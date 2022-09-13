# -*- coding: utf-8 -*-
# @Time    : 2018-10-08 18:58
# @Author  : yao.li
# @Content : 任务

import time
from celery import Celery

# rabbitmq配置
# broker = 'amqp://leeyoung:leeyoung@127.0.0.1/celery/'
# redis配置
broker = 'redis://localhost:6379/1'

# backend = 'amqp://leeyoung:leeyoung@127.0.0.1/celery/'
backend = 'redis://localhost:6379/2'

app = Celery('tasks', broker=broker, backend=backend)


@app.task
def add(x, y):
    print('>>>>>>>>>>>>>>>>> enter called func')
    time.sleep(5)
    print('>>>>>>>>>>>>>>>>> deal end ........')
    return x + y
