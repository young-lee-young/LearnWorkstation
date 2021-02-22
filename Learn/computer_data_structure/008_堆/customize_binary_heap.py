# -*- coding: utf-8 -*-
# @Time:    2020/3/16 8:14 AM
# @Author:  leeyoung
# @File:    binary_heap.py
# @Content: 二叉堆（这里实现的是最大堆）

from random import shuffle


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
    def _sift_down(self, index):
        # 循环结束的条件为节点是叶子节点，即当前节点左孩子节点索引大于数组长度，退出循环
        while self._left_child_index(index) < self.size():
            # 取最大孩子节点索引为左孩子索引
            max_child_index = self._left_child_index(index)
            # 左孩子节点加1为右孩子节点，右孩子节点小于size说明存在右孩子节点；并且右孩子节点大于左孩子节点；则最大孩子节点索引为右孩子节点索引
            if (max_child_index + 1) < self.size() and self.data[max_child_index + 1] > self.data[max_child_index]:
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
        # 从堆最后一个不为叶子节点的节点，往前做slft_down操作
        for index in range(self._parent_index(self.size() - 1), -1, -1):
            self._sift_down(index)


# 生成列表，并打散列表
num_list = list(range(100))
shuffle(num_list)

# 向堆中添加元素
binary_heap = BinaryHeap()
for num in num_list:
    binary_heap.add_node(num)

# 验证堆
new_list = list()
while not binary_heap.is_empty():
    new_list.append(binary_heap.remove_node())
last_num = new_list[0]
for num in new_list[1:]:
    # 最大堆，remove出来的顺序应该是从大到小的，如果上一个值比当前大，说明堆是错误的
    if last_num < num:
        print("---------- error ----------")
        break
    last_num = num

print('---------- is empty ----------', binary_heap.is_empty())
binary_heap.heapify(num_list)
# 验证heapify
new_list = list()
while not binary_heap.is_empty():
    new_list.append(binary_heap.remove_node())
last_num = new_list[0]
for num in new_list[1:]:
    # 最大堆，remove出来的顺序应该是从大到小的，如果上一个值比当前大，说明堆是错误的
    if last_num < num:
        print("---------- error ----------")
        break
    last_num = num
