# 异步IO的学习

import time
import asyncio

now = time.time()


# 定义一个协程函数，协程也是一个对象，不能直接调用，返回一个协程对象
async def get_fun(i):
    print(i)
    await wait()
    return i


def callback(future):
    print(future.result())


def wait():
    time.sleep(10)


coroutine1 = get_fun(1) # coroutine（协程的意思）下面有注册到事件循环，所以这个协程有效
coroutine2 = get_fun(2)
coroutine3 = get_fun(3)


loop = asyncio.get_event_loop() # 创建事件的循环
# 将协程注册到事件循环, run_until_complete参数是一个future对象, 当传入一个协程时, 内部自动封装成task,task是Future的子类
loop.run_until_complete(coroutine1)


# 创建一个task
task = loop.create_task(coroutine2)
print(type(task))
print(task) # <Task pending coro=<get_fun() running at E:/PycharmWorkstation/Code/Learn/001_asyncio.py:14>>
# task.add_done_callback(callback)   #加入回调方法
loop.run_until_complete(task)
print(task) # <Task finished coro=<get_fun() done, defined at E:/PycharmWorkstation/Code/Learn/001_asyncio.py:14> result=None>
# print(task.result())   #协程执行后得到的结果


# 也能创建一个task
task2 = asyncio.ensure_future(coroutine3)
print(type(task2))  # <class '_asyncio.Task'>
loop.run_until_complete(task2)

loop.close()
