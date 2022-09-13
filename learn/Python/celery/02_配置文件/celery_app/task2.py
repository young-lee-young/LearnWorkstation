# -*- coding: utf-8 -*-
# @Time    : 2018-10-08 20:09
# @Author  : yao.li
# @Content :

import time
from celery_app import app


@app.task
def multipy(x, y):
    print('start called func multipy ...')
    time.sleep(3)
    return x * y
