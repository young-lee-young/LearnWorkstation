# 对切片命名

import random

num_list = [random.randint(1,100) for i in range(50)]
print(num_list)

slice_1 = slice(2, 13, 2)    # 开始,结束,步长

new_num_list = num_list[slice_1]
print(new_num_list)

print(slice_1.start)
print(slice_1.stop)
print(slice_1.step)
print('------------------------')