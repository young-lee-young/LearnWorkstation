import random
import time


# 时间复杂度是n,
num_list = [random.randint(1, 100) for _ in range(100)]
print(num_list)


def shell_sort(array):
    gap = len(array)
    while gap > 1:
        gap = gap // 2
        for i in range(gap, len(array)):
            for j in range(i % gap, i, gap):
                if array[i] < array[j]:
                    array[i], array[j] = array[j], array[i]
    return array


new_list = shell_sort(num_list)
print(new_list)
