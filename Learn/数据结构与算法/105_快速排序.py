import random
import time


num_list = [random.randint(-100, 100) for _ in range(100)]
print(num_list)


# 时间复杂度是nlog2n, 控件复杂度nlog2n, 不稳定
def quick_sort(array, left, right):
    if left >= right:   # 如果只有一个数字时，结束递归
        return
    flag = left
    for i in range(left + 1, right + 1):    # 默认以第一个数字作为基准数，从第二个数开始比较，生成索引时要注意右部的值
        if array[flag] > array[i]:
            tmp = array[i]
            del array[i]
            array.insert(flag, tmp)
            flag += 1
    quick_sort(array, left, flag - 1)   # 将基准数前后部分分别递归排序
    quick_sort(array, flag + 1, right)
    return array


new_list = quick_sort(num_list, 0, len(num_list) - 1)
print(new_list)