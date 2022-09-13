# -*- coding: utf-8 -*-


from datetime import timedelta
import djcelery
djcelery.setup_loader()


CELERY_QUEUES = {
    'beat_tasks': {
        'exchange': 'beat_tasks',
        'exchange_type': 'direct',
        'binding_key': 'beat_tasks'
    },
    'work_queue': {
        'exchange': 'work_queue',
        'exchange_type': 'direct',
        'binding_key': 'work_queue'
    }
}

CELERYBEAT_SCHEDULE = {
    'task1': {
        'task': 'test_task',
        'schedule': timedelta(seconds=1),
        'options': {
            'queue': 'beat_tasks'
        }
    }
}

# 默认队列
CELERY_DEFAULT_QUEUE = 'work_queue'

# 导入APP下面的tasks
CELERY_IMPORTS = {
    'celery_test.tasks'
}

# 有些情况下可以防止死锁
CELERYD_FORCE_EXECV = True

# 设置并发worker数量
CELERYD_CONCURRENCY = 4

# 失败重试
CELERY_ACKS_LATE = True

# 每个worker最多执行任务数，然后自动销毁，可以防止内存泄漏
CELERYD_MAX_TASKS_PER_CHILD = 100

# 单个任务最大运行时间，超过会被杀掉
CELERYD_TASK_TIME_LIMIT = 12 * 30
