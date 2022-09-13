# 从序列移除重复的元素,顺序保持不变

import random

num_list = [random.randint(0, 10) for i in range(10)]
print(num_list)


def pop_same(items):
    seen = set()
    for item in items:
        if item not in seen:
            yield item
            seen.add(item)


poped = list(pop_same(num_list))
print(poped)


# 简单的移除重复项,使用集合,移除后改变顺序
poped = set(num_list)
print(poped)
print('------------------------')