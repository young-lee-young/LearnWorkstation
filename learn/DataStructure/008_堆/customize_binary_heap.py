# -*- coding: utf-8 -*-
# @Time:    2020/3/16 8:14 AM
# @Author:  leeyoung
# @File:    binary_heap.py
# @Content: 二叉堆（这里实现的是最大堆）

from random import shuffle
import sys


class BinaryHeap:
    def __init__(self):
        # 使用数组存储数据
        self.data = list()

    def size(self):
        return len(self.data)

    def is_empty(self):
        return True if self.size() == 0 else False

    def top(self):
        return self.data[0]

    # 父节点索引
    def _parent_index(self, index):
        return (index - 1) // 2

    # 左孩子节点
    def _left_child_index(self, index):
        return index * 2 + 1

    # 右孩子节点
    def _right_child_index(self, index):
        return index * 2 + 2

    # 向堆中添加元素，O(logn)，先插入数据，再上浮
    def add_node(self, num):
        self.data.append(num)
        self._sift_up(len(self.data) - 1)

    # 堆元素上浮过程，维持堆性质，堆父节点大于子节点
    def _sift_up(self, index):
        # 新添加的节点数据比父节点大，将两个节点数据交换
        # 父节点是最大的，所以不需要和另一个孩子比较进行sift_down操作
        while index > 0 and self.data[self._parent_index(index)] < self.data[index]:
            parent_index = self._parent_index(index)
            self.data[parent_index], self.data[index] = self.data[index], self.data[parent_index]
            index = parent_index

    # 在堆中删除元素，只能删除堆顶元素，O(logn)
    def remove_node(self):
        head_node = self.data[0]
        # 交换堆顶和最后一个元素
        self.data[0], self.data[-1] = self.data[-1], self.data[0]
        # 删除最大元素
        self.data.pop(-1)
        # 将堆顶下沉
        self._sift_down(0)
        return head_node

    # 堆元素下沉过程，先交换元素，删除最后元素，再下沉
    def _sift_down(self, index, max_index=None):
        # 这里的max_index是为对排序设置的，其他的sift_down设置为数组size即可
        if max_index is None:
            max_index = self.size()
        # 循环结束的条件为节点是叶子节点，即当前节点左孩子节点索引大于数组长度，退出循环
        # while self._left_child_index(index) < self.size():
        while self._left_child_index(index) < max_index:
            '''
            找左右孩子中较大的，获取索引值，当前值比大孩子大，交换当前节点和大孩子
            '''
            # 取左孩子索引为最大孩子节点索引
            max_child_index = self._left_child_index(index)
            # 左孩子节点加1为右孩子节点，右孩子节点小于size说明存在右孩子节点；并且右孩子节点大于左孩子节点；则最大孩子节点索引为右孩子节点索引
            # if (max_child_index + 1) < self.size() and self.data[max_child_index + 1] > self.data[max_child_index]:
            if (max_child_index + 1) < max_index and self.data[max_child_index + 1] > self.data[max_child_index]:
                max_child_index += 1

            # 当前节点比最大孩子大，跳出循环
            if self.data[index] > self.data[max_child_index]:
                break

            # 当前节点和最大孩子节点交换
            self.data[index], self.data[max_child_index] = self.data[max_child_index], self.data[index]
            index = max_child_index

    # 获得最大元素，放入新元素
    def replace(self, num):
        head_node = self.data[0]
        self.data.insert(0, num)
        self._sift_down(0)
        return head_node

    # 将任意数组整理成堆
    # 将元素逐个插入到空堆中，算法复杂度是O(nlogn)，使用heapify算法复杂度O(logn)
    def heapify(self, num_list):
        self.data = num_list
        # 从堆最后一个不为叶子节点的节点，往前做sift_down操作
        #         5                10
        #       /   \             /   \
        #      10    6   --->    8     7
        #     / \   /           / \   /
        #    3   8 7           3   5 6
        # 这里就是从6开始做sift_down操作，然后是10，最后是8
        for index in range(self._parent_index(self.size() - 1), -1, -1):
            self._sift_down(index)

    # 堆排序，将堆heapify后，再将数据一个一个取出
    def heap_sort(self):
        max_index = self.size()
        while max_index > 0:
            # 交换堆顶和最后一个值，将最大值放到最后，再将前面的值排序
            self.data[0], self.data[max_index - 1] = self.data[max_index - 1], self.data[0]
            max_index -= 1
            self._sift_down(0, max_index)


# 验证最大堆
def check_max_heap(binary_heap):
    is_max_heap = True
    # 从堆中删除元素
    last = sys.maxsize
    print('---------- is empty ----------', binary_heap.is_empty())
    while not binary_heap.is_empty():
        # 最大堆，remove出来的顺序应该是从大到小的，如果上一个值比当前大，说明堆是错误的
        num = binary_heap.remove_node()
        if last < num:
            is_max_heap = False
            break
        last = num
    print('---------- is empty ----------', binary_heap.is_empty())
    return is_max_heap


# 生成列表，并打散列表
num_list = list(range(10))
shuffle(num_list)

# 向堆中添加元素
binary_heap = BinaryHeap()
for num in num_list:
    binary_heap.add_node(num)
print('---------- is max heap ----------', check_max_heap(binary_heap))

# 验证heapify
num_list = [5, 10, 6, 3, 8, 7]
binary_heap.heapify(num_list)
print('---------- is max heap ----------', check_max_heap(binary_heap))

# 验证堆排序
binary_heap.heapify(num_list)
binary_heap.heap_sort()
print(binary_heap.data)
