# 找到最大或最小的N个元素

import random
import heapq

num_list = [random.randint(0, 100000) for i in range(10000)]
print(num_list)


num_min = min(num_list)  # 最小的元素
num_max = max(num_list)  # 最大的元素
print(num_min, num_max)


num_lar = heapq.nlargest(2, num_list)  # 最大的两个
num_sma = heapq.nsmallest(2, num_list)  # 最小的两个
print(num_lar, num_sma)


# 在很大的数据里找最大和最小的数据,下面的方法效率高
heapq.heapify(num_list)  # 元素以堆的形式排列,最小的元素在第一位
print(num_list)


num_pop = heapq.heappop(num_list)   # 从开始去除一个元素
print(num_pop)


print(num_list) # 序列改变
print('------------------------')