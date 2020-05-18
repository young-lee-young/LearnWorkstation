import random
from datetime import datetime


# FIND_NUM = 888888
FIND_NUM = 1


def time_counter(func):
    def inner_func(num_list=None):
        print('--------------- function_nam: {}'.format(func.__name__))
        start_time = datetime.now()
        num_list = func(num_list) if num_list else func()
        end_time = datetime.now()
        print('time consuming: {}'.format(end_time - start_time))
        return num_list
    return inner_func


# 顺序查找
@time_counter
def sequence_search(num_list):
    # 查找成功最少比较 1 次，最多比较 n 次，平均(1 + 2 + 3 + ... + n) / n = (n * (n + 1) / 2) / n = (n + 1) / 2
    # 查找失败比较次数: n 次

    # 时间复杂度：O(n)
    for index, num in enumerate(num_list):
        if num == FIND_NUM:
            return index
    return -1


# 二分查找
@time_counter
def binary_search(num_list):
    # 每次查找区间大小 n，n / 2，n / 4，...，n / 2^k，k为循环次数，n / 2^k = 1，k = logn

    # 时间复杂度：O(logn)
    low_index = 0
    high_index = len(num_list) - 1

    while low_index <= high_index:
        middle_index = int(low_index + (high_index - low_index) * 1 / 2)
        if FIND_NUM > num_list[middle_index]:
            low_index = middle_index + 1
        elif FIND_NUM < num_list[middle_index]:
            high_index = middle_index - 1
        else:
            return middle_index
    return -1


# 插值查找
@time_counter
def insert_search(num_list):
    # 时间复杂度 log(logn)
    low_index = 0
    high_index = len(num_list) - 1

    while low_index <= high_index:
        middle_index = int(low_index + (high_index - low_index) * (FIND_NUM - num_list[low_index]) / (num_list[high_index] - num_list[low_index]))
        if FIND_NUM > num_list[middle_index]:
            low_index = middle_index + 1
        elif FIND_NUM < num_list[middle_index]:
            high_index = middle_index - 1
        else:
            return middle_index
    return -1


# 插入排序 - 直接插入排序
def insert_sort(num_list):
    # 最好情况：外层循环 n - 1 次，每次进行 1 次比较，一共比较 (n - 1) 次，移动 0 次
    # 时间复杂度：O(1)

    # 最坏情况：外层循环 n - 1 次，每次进行 1 + 2 + ... + (n - 2) + (n - 1) = n(n - 1) / 2 次比较，移动(1 + 2) + (2 + 2) + ... + (n - 2 + 2) + (n - 1 + 2) = (n + 4)(n - 1) / 2
    # 时间复杂度：O(n^2)
    num_list_len = len(num_list)

    for i in range(1, num_list_len):
        for j in reversed(range(i)):
            if num_list[i] < num_list[j]:
                num_list[j], num_list[i] = num_list[i] = num_list[j]
                i = j
            else:
                break
    return num_list


# 插入排序 - 折半插入排序
def half_insert_sort(num_list):
    # 时间复杂度O(n^2)
    num_list_len = len(num_list)

    for i in range(1, num_list_len + 1):
        for j in reversed(range(i)):
            # 二分查找代码
            low_index = 0
            high_index = j
            middle_index = int(low_index + (high_index - low_index) * 1 / 2)
            while low_index <= high_index:
                middle_index = int(low_index + (high_index - low_index) * 1 / 2)
                if num_list[j] > num_list[middle_index]:
                    low_index = middle_index + 1
                elif num_list[j] < num_list[middle_index]:
                    high_index = middle_index - 1
                else:
                    break

            # 移动元素
            for k in range(j, middle_index, -1):
                num_list[k - 1], num_list[k] = num_list[k], num_list[k - 1]
    return num_list


# 插入排序 - 希尔排序
def shell_sort(num_list):
    # Hibbard增量序列（增量是2^k - 1），最坏情况 O(n^1.5)，平均情况 O(n^1.25)
    # Sedgewick增量序列，最坏情况 O(n^1.33...)，平均情况：O(n^(7/6))
    num_list_len = len(num_list)

    gap = num_list_len
    while gap > 1:
        gap = gap // 2
        for i in range(gap, num_list_len):
            for j in range(i % gap, i, gap):
                if num_list[i] < num_list[j]:
                    num_list[i], num_list[j] = num_list[j], num_list[i]
    return num_list


# 交换排序 - 冒泡排序
@time_counter
def bubble_sort(num_list):
    # 最好情况：比较次数：n - 1 次，移动次数：0
    # 最好情况：循环次数：(n - 1) + (n - 2) + (n - 3) + ... + 2 + 1 = n(n - 1) / 2，移动次数：3n(n - 1) / 2

    # 时间复杂度O(n^2)
    num_list_len = len(num_list)

    for i in range(num_list_len - 1):
        for j in range(num_list_len - 1):
            if num_list[j] > num_list[j + 1]:
                num_list[j], num_list[j + 1] = num_list[j + 1], num_list[j]
        num_list_len = num_list_len - 1
    return num_list


# 交换排序 - 快速排序
@time_counter
def quick_sort(num_list):
    pass


# 选择排序 - 简单选择排序
@time_counter
def select_sort(num_list):
    # 最好情况：比较次数：1 + 2 + ... + (n - 1) = n(n - 1) / 2，移动次数0次
    # 最坏情况：比较次数：1 + 2 + ... + (n - 1) = n(n - 1) / 2，移动次数

    # 时间复杂度：O(n^2)
    num_list_len = len(num_list)

    for index in range(num_list_len - 1):
        min_index = index
        for j in range(index + 1, num_list_len):
            if num_list[min_index] > num_list[j]:
                min_index = j
        num_list[index], num_list[min_index] = num_list[min_index], num_list[index]
    return num_list


# 选择排序 - 堆排序
@time_counter
def heap_sort():
    pass


@time_counter
def generate_num_list():
    num_list = [random.randint(1, 50000) for i in range(2000)]
    num_list.append(FIND_NUM)
    num_list = list(set(num_list))
    print('num list len: {}'.format(len(num_list)))
    return num_list


def main(sort=None, find=None):
    num_list = generate_num_list()

    sort_func_dict = {
        'bubble': bubble_sort
    }
    num_list = sort_func_dict.get(sort)(num_list) if sort else num_list
    print(num_list)

    find_func_dict = {
        'sequential': sequence_search,
        'binary': binary_search,
        'insert': insert_search
    }
    index = find_func_dict.get(find)(num_list)
    print('--------------- index: {}'.format(index))
    print(num_list[index - 5: index + 5])


if __name__ == '__main__':
    sort_dict = {
        '冒泡排序': 'bubble',
        '快速排序': 'quick'
    }
    finc_dict = {
        '顺序查找': 'sequence',
        '二分查找': 'binary',
        '插值查找': 'insert'
    }
    sort = sort_dict.get('冒泡排序')
    find = finc_dict.get('插值查找')

    main(sort=sort, find=find)
