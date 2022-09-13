# -*- coding: utf-8 -*-
# @Time    : 2018-10-08 18:52
# @Author  : yao.li
# @Content : celery调度函数

from tasks import add


if __name__ == '__main__':
    print('start task ...')
    while True:
        result = add.delay(2, 8)
    # 定义队列的名字
    # result = add.apply_async((2, 8), queue='celery_add')
        print('end task ...')
        print(result.id)
        print(result.status)
