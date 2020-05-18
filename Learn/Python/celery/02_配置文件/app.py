# -*- coding: utf-8 -*-
# @Time    : 2018-10-08 20:16
# @Author  : yao.li
# @Content :


from celery_app import task1, task2


print('start ...')
# task1.add.apply_async 会报 'task args must be a list or tuple'错误
task1.add.delay(2, 4)
task2.multipy.delay(2, 4)
print('end ...')
