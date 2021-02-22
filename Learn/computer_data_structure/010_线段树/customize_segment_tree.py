# -*- coding: utf-8 -*-
# @Time:    2020/3/16 5:49 PM
# @Author:  leeyoung
# @File:    segment_tree.py
# @Content: 线段树（区间树）


# 线段树常用于区间查询，并且数组不变
class SegmentTree:
    def __init__(self, num_list):
        # 使用数组来存储元素
        self.data = num_list
        # 有n个元素，需要使用4n的空间来存储
        self.tree = [0] * (4 * len(self.data))
        self._build_segment_tree(0, 0, len(self.data) - 1)

    # 构建线段树，index表示表示创建区间[begin, end]的线段树
    def _build_segment_tree(self, index, begin, end):
        # 只剩一个元素，已经到叶子结点
        if begin == end:
            self.tree[index] = self.data[begin]
            return

        # 左、右孩子索引
        left_child_index = self._left_child_index(index)
        right_child_index = self._right_child_index(index)
        # (begin + end) // 2可能产生整型溢出
        mid = begin + (end - begin) // 2

        self._build_segment_tree(left_child_index, begin, mid)
        self._build_segment_tree(right_child_index, mid + 1, end)

        # 求和线段树：当前线段树 = 左线段树 + 右线段树
        self.tree[index] = self.tree[left_child_index] + self.tree[right_child_index]

    # 返回左孩子节点索引
    def _left_child_index(self, index):
        return index * 2 + 1

    # 返回右孩子节点索引
    def _right_child_index(self, index):
        return index * 2 + 2

    def size(self):
        return len(self.data)

    def get(self, index):
        if index < 0 or index >= len(self.data):
            return None
        if self.size() == 0 or self.data is None:
            return None
        return self.data[index]

    # 查询区间，O(logn)
    def query(self, begin, end):
        return self._query(0, 0, self.size() - 1, begin, end)

    # 在以index为根的线段树中[begin, end]范围里，搜索区间[query_begin, query_end]的值
    def _query(self, index, begin, end, query_begin, query_end):
        if begin == query_begin and end == query_end:
            return self.tree[index]

        # 左、右孩子节点值
        left_child_index = self._left_child_index(index)
        right_child_index = self._right_child_index(index)

        # mid将当前线段树分成左右两部分
        mid = begin + (end - begin) // 2

        # 查询左边界比中间值大，说明查询区间在右子树里
        if query_begin >= mid + 1:
            return self._query(right_child_index, mid + 1, end, query_begin, query_end)
        # 查询右边界比中间值小，说名查询区间在左子树里
        elif query_end <= mid:
            return self._query(left_child_index, begin, mid, query_begin, query_end)
        # 查询区间即在左子树又在右子树里
        else:
            # 查找左边
            left_result = self._query(left_child_index, begin, mid, query_begin, mid)
            # 查找右边
            right_result = self._query(right_child_index, mid + 1, end, mid + 1, query_end)
            return left_result + right_result

    # 更新元素，O(logn)
    def set(self, index, value):
        self.data[index] = value
        self._set(0, 0, self.size() - 1, index, value)

    # 将index为根的线段树中的index更新为value
    def _set(self, index, begin, end, set_index, value):
        if begin == end:
            self.tree[index] = value
            return

        # 左右孩子节点值
        left_child_index = self._left_child_index(index)
        right_child_index = self._right_child_index(index)

        # mid将线段树为分左右两部分
        mid = begin + (end - begin) // 2

        # 如果需要更新的节点大于中间值，说明更新节点在右子树
        if set_index >= mid + 1:
            self._set(right_child_index, mid + 1, end, set_index, value)
        # 需要更新的节点小于中间值，说明更新节点在左子树
        else:
            self._set(left_child_index, begin, mid, set_index, value)

        # 更新节点的父节点均受到了影响，所以需要重新计算
        self.tree[index] = self.tree[left_child_index] + self.tree[right_child_index]


# 构造测试数据
num_list = [-2, 0, 3, -5, 2, -1]
segment_tree = SegmentTree(num_list)

# 测试查询方法
print('---------- query ----------', segment_tree.query(0, 2))

# 测试更新方法
print('---------- set ----------', segment_tree.set(1, 2))
print('---------- query ----------', segment_tree.query(0, 2))
