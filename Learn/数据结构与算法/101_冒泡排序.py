import random
import time


# num_list = [random.randint(1, 100) for _ in range(100)]
# print('原数据是',num_list)
#
#
# def bubble_sort(array):
#     change_times = 0
#     for i in range(len(array)):
#         for j in range(i+1, len(array)):
#             # 每一步的目的就是把最小的最前面
#             if array[i] > array[j]:
#                 # 下面这写法Python是成立的
#                 array[i], array[j] = array[j], array[i]
#                 change_times += 1
#     print(change_times)
#     return array
#
#
# new_list = bubble_sort(num_list)
# # 此时num_list是排序过的
# print(num_list)


# 两个相邻的比较，如果前一个比第二个大，两个调换位置，每次都是把最大的放到最后面
# 时间复杂度是n^2, 控件复杂度是1, 稳定
num_list2 = [random.randint(1, 100) for _ in range(100)]
print(num_list2)


def bubble_sort2(array):
    change_times = 0
    for i in range(len(array) - 1):
        for j in range(len(array) - i - 1):
            if array[j] > array[j+1]:
                array[j], array[j+1] = array[j+1], array[j]
                change_times += 1
    print(change_times)
    return array


new_list2 = bubble_sort2(num_list2)
print(new_list2)
