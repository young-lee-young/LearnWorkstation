# 序列中出现次数最多的元素

import random
from collections import Counter

num_list = [random.randint(1, 5) for i in range(10)]

num_count = Counter(num_list)
print(num_count)

top_three = num_count.most_common(2)
print(top_three)

more_num = [random.randint(1, 5) for i in range(5)]

for num in more_num:
    num_count[num] += 1
# num_count.update(more_num)    # 上面三句还可以用此语句代替
print(num_count)
print('------------------------')
