# 让带有N个参数的可调用对象以比较少的参数形式调用

from functools import partial

def fun_1(a, b, c, d):
    print(a, b, c, d)

s1 = partial(fun_1, 1)  # a=1
print(s1(2, 3, 4))

s2 = partial(fun_1, d=20)   # d=20
print(s2(1, 2, 3))

s3 = partial(fun_1, 1, 2, d=20) # a=1, b=2, d=20
print(s3(3))
print('----------------------')


# 按点之间的距离排序
import math

points = [(1, 2), (3, 4), (5, 6), (7, 8)]

def distance(point1, point2):
    x1, y1 = point1
    x2, y2 = point2
    return math.hypot(x2 - x1, y2 - y1)

point = (3, 5)
points.sort(key=(partial(distance, point)))
print(points)
print('----------------------')


import logging
from multiprocessing import Pool

def output_result(result, log=None):
    if log is not None:
        log.debug('Got %r', result)

def add(a, b):
    return a + b

if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    log = logging.getLogger('test')

    """
    这个里有一个不理解的地方,走到这一步,上面的函数会被执行4次
    """
    p = Pool()
    p.apply_async(add, (3, 4),callback=partial(output_result, log=log))
    p.close()
    p.join()