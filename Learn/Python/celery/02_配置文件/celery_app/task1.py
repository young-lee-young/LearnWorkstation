# -*- coding: utf-8 -*-
# @Time    : 2018-10-08 20:09
# @Author  : yao.li
# @Content :

import time
from celery_app import app


@app.task
def add(x, y):
    print('>>>>>>>>>>>>>> start called func add')
    time.sleep(3)
    print('>>>>>>>>>>>>>> deal end')
    return x + y
