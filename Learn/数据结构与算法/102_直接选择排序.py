import random
import time


num_list = [random.randint(1, 100) for _ in range(100)]
print(num_list)


# 时间复杂度是n^2, 空间复杂度是1, 不稳定
def selection_sort(array):
    for i in range(0, len(array)):
        min = i
        for j in range(i + 1, len(array)):
            # 如果最小的比找到的大，就把找到的数小标记为min
            if array[min] > array[j]:
                min = j
        # 把找到的最小的那个和第i位调换
        array[i], array[min] = array[min], array[i]
    return array


new_list = selection_sort(num_list)
print(new_list)
