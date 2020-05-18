import random
import time


num_list = [random.randint(1, 100) for _ in range(100)]
print(num_list)


# 默认第一个数字有顺序，每次取一个数，与前面的比较，插入到相应的位置
# 时间复杂度是n^2, 空间复杂度是1, 稳定
def insertion_sort(array):
    # i即将插入到已排好序的数列中的数
    for i in range(1, len(array)):
        # 已经排好序的数列
        for j in range(i):
            if array[i] < array[j]:
                array.insert(j, array.pop(i))
                break
    return array


new_list = insertion_sort(num_list)
print(new_list)


