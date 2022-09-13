# -*- coding: utf-8 -*-
# @Time:    2020/3/14 7:48 PM
# @Author:  leeyoung
# @File:    binary_search_tree.py
# @Content: 二分搜索树

from collections import deque
from random import shuffle
from tree_node import BinTreeNode


# 没有解决平衡的问题，有可能退化为链表
class BinarySearchTree:
    def __init__(self):
        self.root = None
        self.tree_size = 0

    def size(self):
        return self.tree_size

    def is_empty(self):
        return True if self.size() == 0 else False

    # 添加一个节点
    def add_node(self, num):
        if not self.root:
            self.root = BinTreeNode(num)
            self.tree_size += 1
        else:
            self._add_node(self.root, num)

    def _add_node(self, node, num):
        if node.data == num:
            return
        # 值小于当前节点值，并且当前节点没有左子树节点
        if num < node.data and not node.left:
            node.left = BinTreeNode(num)
            self.tree_size += 1
            return
        # 值大于当前节点值，并且当前节点没有右子树节点
        if num > node.data and not node.right:
            node.right = BinTreeNode(num)
            self.tree_size += 1
            return

        # 往左子树添加
        if num < node.data:
            self._add_node(node.left, num)
        # 往右子树添加
        else:
            self._add_node(node.right, num)

    # 添加一个节点，新方法
    def add_node_new(self, num):
        self.root = self._add_node_new(self.root, num)

    def _add_node_new(self, node, num):
        if not node:
            self.tree_size += 1
            node = BinTreeNode(num)
            return node

        if num < node.data:
            node.left = self._add_node_new(node.left, num)
        if num > node.data:
            node.right = self._add_node_new(node.right, num)

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

    # 先序遍历
    def preorder_traversal(self):
        self._preorder_traversal(self.root)

    def _preorder_traversal(self, node):
        if node is not None:
            print(node.data)
            self._preorder_traversal(node.left)
            self._preorder_traversal(node.right)

    # 先序遍历，非递归实现（使用栈实现）
    def preorder(self):
        stack = deque()
        stack.append(self.root)
        while len(stack) != 0:
            cur_node = stack.pop()
            print(cur_node.data)

            if cur_node.right:
                stack.append(cur_node.right)
            if cur_node.left:
                stack.append(cur_node.left)

    # 中序遍历
    def inorder_traversal(self):
        self._inorder_traversal(self.root)

    def _inorder_traversal(self, node):
        if node is not None:
            self._inorder_traversal(node.left)
            print(node.data)
            self._inorder_traversal(node.right)

    # 后序遍历
    def postorder_traversal(self):
        self._postorder_traversal(self.root)

    def _postorder_traversal(self, node):
        if node is not None:
            self._postorder_traversal(node.left)
            self._postorder_traversal(node.right)
            print(node.data)

    # 层序遍历，使用队列实现
    def sequence_traversal(self):
        queue = deque()
        queue.append(self.root)

        while queue:
            node = queue.popleft()
            print(node.data)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)

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

    # 获取最大元素
    def maximum(self):
        if self.tree_size == 0:
            return None
        return self._maximum(self.root).data

    # 最大值在最右边
    def _maximum(self, node):
        while node.right:
            node = node.right
        return node

    # 删除最小节点
    def remove_min(self):
        node = self.minimum()
        self._remove_min(self.root)
        return node

    def _remove_min(self, node):
        if not node.left:
            right_node = node.right
            node.right = None
            self.tree_size -= 1
            return right_node

        node.left = self._remove_min(node.left)
        return node

    # 删除最大节点
    def remove_max(self):
        node = self.maximum()
        self._remove_max(self.root)
        return node

    def _remove_max(self, node):
        if not node.right:
            left_node = node.left
            node.left = None
            self.tree_size -= 1
            return left_node

        node.right = self._remove_max(node.right)
        return node

    # 删除任意节点
    def remove_node(self, num):
        self.root = self._remove_node(self.root, num)

    def _remove_node(self, node, num):
        if not node:
            return
        if num < node.data:
            node.left = self._remove_node(node.left, num)
            return node
        elif num > node.data:
            node.right = self._remove_node(node.right, num)
            return node
        else:
            # 只有右子树
            if not node.left:
                right_node = node.right
                node.right = None
                self.tree_size -= 1
                return right_node
            # 只有左子树
            if not node.right:
                left_node = node.left
                node.left = None
                self.tree_size -= 1
                return left_node
            # 既有左子树，又有右子树
            successor = self._minimum(node.right)
            successor.right = self._remove_min(node.right)
            successor.left = node.left
            # 删除node节点
            node.left = node.right = None
            return successor


binary_search_tree = BinarySearchTree()
num_list = list(range(100))
shuffle(num_list)
for num in num_list:
    binary_search_tree.add_node(num)

binary_search_tree_new = BinarySearchTree()
shuffle(num_list)
for num in num_list:
    binary_search_tree_new.add_node_new(num)

# 包含元素
print(binary_search_tree_new.contains(10))
print(binary_search_tree_new.contains(101))

# 递归遍历
binary_search_tree_new.preorder_traversal()
binary_search_tree_new.inorder_traversal()
binary_search_tree_new.postorder_traversal()

# 非递归遍历
binary_search_tree_new.preorder()

# 层序遍历
binary_search_tree_new.sequence_traversal()

# 删除最小值
print(binary_search_tree_new.minimum())
binary_search_tree_new.remove_min()

# 删除最大值
print(binary_search_tree_new.maximum())
binary_search_tree_new.remove_max()

# 删除任意值
binary_search_tree_new.remove_node(98)
print(binary_search_tree_new)
