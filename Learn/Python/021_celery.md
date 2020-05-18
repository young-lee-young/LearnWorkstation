
# celery学习

### 基础知识

1. Brokers

中间人，任务队列本身，rabbitmq, redis, zookeeper

2. backend

存储结果，队列中的任务完成的结果或者状态需要被任务的发送者知道

3. workers

类似消费模型中的消费者，从队列中取出任务并执行

4. tasks



### 使用

1. 耗时任务文件

```python
from celery import Celery

broker='redis://localhost:6379/0'
backend='redis://localhost:6379/1'
app = Celery('tasks', broker=broker, backend=backend)

@app.task
def add(x, y):
	print(x + y)
```

2. 调度文件

```python
result = add.delay(2, 8)	# 需要调用delay，
while not result.ready():
	time.sleep(1)

# delay返回的是一个AsyncResult对象，存的是一个异步的结果，任务完成时result.ready()返回true，使用result.get()获取结果
```


### Django中使用celery

2. 项目中配置

```python
# 在settings同级新建文件celeryconfig.py
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

# 默认的队列
CELERY_DEFAULT_QUEUE = 'work_queue'

CELERY_IMPORTS = (
    'app.任务文件'
)

# 有些情况下可以防止死锁
CELERYD_FORCE_EXECV = True

# 并发worker数，根据CPU数设置
CELERYD_CONCURRENCY = 4

# 允许重试
CELERY_ACKS_LATE = True

# 每个worker最多执行100个任务被销毁，可以防止内存泄漏
CELERYD_MAX_TASKS_PER_CHILD = 100

# 单个任务的最大运行时间
CELERYD_TASK_TIME_LIMIT = 12 * 30
```

```python
# 耗时任务文件
import time
from celery.task import Task


class TimeTask(Task):
    name = 'time-task'
    def run(self, *args, **kwargs):
        time.sleep(5)
        print('args={}, kwargs={}'.format(args, kwargs))
```

3. 将Celery和Django关联

```python
# settings配置文件中
# 在INSTALLED_APPS中增加'djcelery'

from .celeryconfig import *
BROKER_BACKEND = 'redis'
BROKER_URL = 'redis://localhost:6379/1'
CELERY_RESULT_BACKEND = 'redis://localhost:6379/2'
```

4. 启动

```
python manage.py celery worker -l INFO
```
