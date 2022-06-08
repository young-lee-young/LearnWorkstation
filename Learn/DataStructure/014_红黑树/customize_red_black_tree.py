# -*- coding: utf-8 -*-
# @Time:    2021/2/1 上午11:44
# @Author:  leeyoung
# @File:    customize_red_black_tree.py
# @Content: 红黑树

"""
红黑树特点
1. 具有二分搜索树的特点
2. 根节点是黑色的
3. 每个叶子节点都是黑色的空节点
4. 如果一个节点是红色的，那么他的孩子节点都是黑色的
5. 从任意节点到叶子节点的所有路径包含相同数目的黑节点

2-3树
1. 满足二分搜索树的基本性质
2. 节点可以存放一个元素或者两个元素
3. 每个节点有2个或者3个孩子（有2个孩子的节点称为2节点，有3个孩子的节点称为3节点）
4. 绝对平衡（根节点到任意叶子节点所经过节点一定相等）
"""


from random import shuffle
from tree_node import RedBlackTreeNode


# 解决AVL树的插入、删除操作多次再平衡问题
# 红黑树是一种不严格的平衡树，是"黑平衡"二叉树，最大高度为2logn，添加、删除、查找操作时间复杂度为O(logn)
# 红色节点向左倾斜
class RedBlackTree:
    def __init__(self):
        self.root = None
        self.tree_size = 0

    def size(self):
        return self.tree_size

    def is_empty(self):
        return True if self.size() == 0 else False

    # 判断一个节点是否是红色
    def is_red(self, node):
        # 红黑树空节点为黑色
        if node is None:
            return False
        return node.color

    # 左旋转，对node节点进行左旋转，返回旋转后的新根节点x
    #    node                    x
    #    / \                    / \
    #   T1  x    - - - - ->   node  T3
    #      / \                / \
    #     T2 T3              T1 T2
    # 对应：node没有孩子，向node右侧添加元素
    def _left_rotate(self, node):
        # 缓存右子树x
        new_root = node.right

        # 左旋转
        node.right = new_root.left
        new_root.left = node

        # 将新根节点的颜色设置为和旧根节点颜色相同
        # 注意：如果原来的node为红色，这里x和node都为红色
        # 左旋转可能生成连续红色节点，在后续再处理
        new_root.color = node.color
        # 设置旧根节点的颜色为红色，node和x相当于2-3树中的3节点
        node.color = True

        return new_root

    # 右旋转，对node节点进行右旋转，返回旋转后的新节点x
    #    node                  x
    #    / \                  / \
    #   x  T2   - - - - ->   y node
    #  / \                     / \
    # y  T1                   T1 T2
    # 对应：node有红色左孩子，向node左孩子左侧添加元素
    def _right_rotate(self, node):
        # 缓存左子树
        new_root = node.left

        # 左旋转
        node.left = new_root.right
        new_root.right = node

        # 将新根节点的颜色设置为和旧根节点颜色相同
        new_root.color = node.color
        # 将旧根节点的颜色设置为红色，表示和父节点融合在一起
        node.color = True

        return new_root

    # 颜色翻转
    # 对应：node有红色左孩子，向node右侧添加元素
    def _flip_colors(self, node):
        # 将node设置为红色
        node.color = True
        # 将node左右孩子设置为黑色
        node.left.color = False
        node.right.color = False

    # 添加一个节点，递归方法
    def add_node(self, num):
        self.root = self._add_node(self.root, num)
        # 红黑树根为黑色，对根节点进行重新着色
        self.root.color = False

    def _add_node(self, node, num):
        if not node:
            self.tree_size += 1
            # 默认添加红色节点
            node = RedBlackTreeNode(num)
            return node

        if num < node.data:
            node.left = self._add_node(node.left, num)
        if num > node.data:
            node.right = self._add_node(node.right, num)

        # ---------------- 以上为平衡树常规添加节点操作，以下为红黑树左旋转、右旋转、颜色翻转等操作 ------------------
        #    黑           黑                黑              黑                 红
        #   /           /     红1左旋转    /    黑右旋转    /   \   颜色翻转    /   \
        # 红   - - ->  红1     - - ->    红2    - - ->   红    红  - - ->    黑   黑
        #               \              /
        #                红2          红1
        #  1             2               3                 4                  5
        #
        # 1中添加比红色节点大的元素形成2
        # 1中添加比红色节点小的元素直接形成3
        # 1中添加比黑色节点大的元素直接形成4

        # 判断是否需要左旋转
        if self.is_red(node.right) and not self.is_red(node.left):
            node = self._left_rotate(node)

        # 判断是否需要右旋转
        if self.is_red(node.left) and self.is_red(node.left.left):
            node = self._right_rotate(node)

        # 判断是否需要颜色翻转
        if self.is_red(node.left) and self.is_red(node.right):
            self._flip_colors(node)

        return node

    # 查询元素
    def contains(self, num):
        return self._contains(self.root, num)

    def _contains(self, node, num):
        if not node:
            return False
        elif num == node.data:
            return True
        elif num < node.data:
            return self._contains(node.left, num)
        else:
            return self._contains(node.right, num)

    # 获取最小元素
    def minimum(self):
        if self.tree_size == 0:
            return None
        return self._minimum(self.root).data

    # 最小值在最左边
    def _minimum(self, node):
        while node.left:
            node = node.left
        return node

    # 删除任意节点，比较复杂，暂时不实现
    def remove_node(self, num):
        pass


red_black_tree = RedBlackTree()
num_list = list(range(10))
shuffle(num_list)

# 添加元素
for num in num_list:
    red_black_tree.add_node(num)

# 包含元素
print(red_black_tree.contains(1))
print(red_black_tree.contains(10))
