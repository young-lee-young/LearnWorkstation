# -*- coding: utf-8 -*-

import time
from celery.task import Task


class TestTask(Task):
    name = 'test_task'

    def run(self, *args, **kwargs):
        print('>>>>>>>>>> start task')
        time.sleep(5)
        print('args={}, kwargs={}'.format(args, kwargs))
        print('>>>>>>>>>> end task')
